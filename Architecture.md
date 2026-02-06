# Architecture and Reconciliation Logic

## Overview

The system monitors network port bindings using eBPF and synchronizes this state with a userspace Database (BadgerDB). It relies on two main sources of information:

1.  **eBPF Ring Buffer (Events):** Real-time stream of kernel events.
    *   **Open Events:** Captured when a process binds to a port (e.g., `bind` or `listen` syscalls).
    *   **Close Events:** Captured when a socket is closed (e.g., `close` syscall).
2.  **Database (State):** The persistent store representing the last known "truth" of which ports are currently held by which processes.

## Reconciliation Process

To ensure consistency, we process events from the Ring Buffer and compare them against the current state in the Database. For any given port `P`, we look at three signals to decide the next state:

1.  **DB State (DB):** Is the port currently recorded in the database? (`1` = Yes, `0` = No)
2.  **Open Event (RB):** Did we receive an *Open* event for this port from the Ring Buffer? (`1` = Yes, `0` = No)
3.  **Close Event (RB):** Did we receive a *Close* event for this port from the Ring Buffer? (`1` = Yes, `0` = No)

## Truth Table

The following table serves as the "Table of Truth" for determining the necessary action based on these three inputs.

| DB (Current State) | Open Event (From RB) | Close Event (From RB) | Meaning | Action to take |
| :---: | :---: | :---: | :--- | :--- |
| **0** | **0** | **0** | Port never seen | **Nothing** |
| **0** | **1** | **0** | New port binding detected | **Insert to DB** |
| **0** | **0** | **1** | Close event received but port not in DB | **Ignore** (We likely missed the open event or it happened before we started) |
| **0** | **1** | **1** | Port was opened and quickly closed | **Don't insert** (Transient/Short-lived connection) |
| **1** | **0** | **0** | Stable port, no new events | **Keep in DB** |
| **1** | **1** | **0** | Re-bind on an existing port (potentially new PID) | **Update DB** with new PID/Process info |
| **1** | **0** | **1** | Existing port was closed | **Delete from DB** |
| **1** | **1** | **1** | Port closed and immediately successfully reopened | **Update DB** with new binding info |
