# Spiral 
UOR DPL Runtime

## Introduction
The Spiral UOR DPL Runtime client executes WebAssembly Modules (WASM) that implement the UOR DPL Runtime interface. The UOR DPL Runtime interface defines a
set of WASM functions that can be used to interact with a UOR formatted knowledge graph.

## Commands

1. Spiral help - Prints help information for the Spiral client.
2. Spiral run - Runs the UOR DPL Runtime Interface Execute Entrypoint
3. Spiral write - Runs the UOR DPL Runtime Interface Write Entrypoint
4. spiral read - Runs the UOR DPL Runtime Interface Read Entrypoint 
6. spiral delete - Runs the UOR DPL Runtime Interface Delete Entrypoint

## Explanation

A UOR knowledge graph is composed of nodes and edges that are both the same type, a UOR Element. UOR Elements are identified by their relationships within the UOR knowledge graph. Once a UOR Element is identified, its locator and resource type information are used to identify the WASM that processes that UOR Element's represented Class.

