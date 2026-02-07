# guess-it-1

## Overview

**Language:** Go (Golang)

**guess-it-1** is a stateful number prediction program built using a **Test Driven Development (TDD)** approach.

The program reads a sequence of integers from standard input and, **before** the next number is received, outputs a numeric range in which the next value is expected to fall.

Scoring depends on accuracy and precision: the smaller the correct range, the better the score.

---

## Development Methodology (TDD)

This project follows **Test Driven Development**:

1. Tests define the expected behavior of each stage.
2. Logic is implemented only to satisfy failing tests.
3. Refactoring is done continuously without breaking behavior.

TDD ensures:

* Predictable behavior under large datasets
* Safe refactoring of prediction strategies
* Clear separation of responsibilities

---

## How It Works

1. The program reads **one integer at a time** from stdin.
2. It keeps an internal state of all previous values.
3. Statistical calculations are updated incrementally.
4. A prediction range for the **next** number is generated.
5. The range is printed in the format:

   ```
   lower upper
   ```

---

## Architecture

The implementation follows a **pipeline architecture**, where each stage performs a single, well-defined task.

```
stdin
 ↓
Input Parser
 ↓
State Manager
 ↓
Statistics Engine
 ↓
Prediction Strategy
 ↓
stdout (lower upper)
```

### Pipeline Stages

#### 1. Input Parser

* Reads the current value from stdin
* Validates and converts input to integer

#### 2. State Manager

* Stores historical values
* Exposes state immutably to downstream stages

#### 3. Statistics Engine

Incrementally computes:

* Mean
* Standard Deviation
* Minimum / Maximum
* Trend based on recent values

#### 4. Prediction Strategy

* Combines statistical indicators
* Computes a probabilistic prediction window
* Balances precision against safety

Example logic:

```
center = mean + trend
lower  = center - k * std
upper  = center + k * std
```

#### 5. Output Stage

* Prints the prediction range to stdout
* Contains no internal state

---

## Prediction Strategy Philosophy

The strategy is **probabilistic, not deterministic**.

Characteristics:

* Adapts dynamically to data behavior
* Responds to volatility spikes
* Shrinks the prediction window as confidence grows

The strategy can be replaced without affecting the rest of the system.

---

## Performance Considerations

* O(1) operations per input value
* Incremental statistics (no full recomputation)
* Suitable for large datasets used by automated testers

---

## Running With the Tester

Required directory structure:

```
root/
 ├─ student/
 │   ├─ solution.(go | js | py | rs)
 │   └─ README.md
 └─ script.sh
```

Example `script.sh`:

```sh
#!/bin/sh
node ./student/solution.js
```

Make sure `script.sh` is executable.

---

## Summary

**guess-it-1** is a statistical forecasting exercise focused on:

* clean pipeline architecture
* TDD-driven design
* efficient stateful computation

There is no perfect prediction — only well-structured reasoning supported by data.
