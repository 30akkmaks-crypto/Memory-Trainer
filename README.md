🧠 Memory Trainer (Numbers) — Multi‑Language Brain Exercise
8 languages, one powerful memory trainer – test and improve your short‑term memory by recalling sequences of numbers with increasing difficulty.

✨ Features
🔢 Random number sequences – from 0 to 9, adjustable length

📈 Multiple difficulty levels – easy (4 digits), medium (6), hard (8+)

⏱️ Timed display – numbers shown for a configurable duration

✅ Instant feedback – check your answer and see score per round

🎯 Scoring system – track correct answers across multiple rounds

🖥️ Cross‑platform CLI – works on any terminal

🧰 Supported Languages & Files
Language	File
Python	memory_trainer.py
Go	memory_trainer.go
JavaScript (Node)	memory_trainer.js
Ruby	memory_trainer.rb
PHP	memory_trainer.php
Java	MemoryTrainer.java
C#	MemoryTrainer.cs
C++	memory_trainer.cpp
🚀 Common Usage
All implementations follow the same CLI pattern:

bash
# Start with default settings (medium, 1 round)
<command>

# Specify number of digits (length)
<command> --length 5

# Specify difficulty: easy (4), medium (6), hard (8)
<command> --difficulty hard

# Play multiple rounds
<command> --rounds 3 --length 7

# Show help
<command> --help
Arguments:

-l, --length <n> – number of digits per sequence (default: 6)

-d, --difficulty <easy|medium|hard> – preset lengths (4, 6, 8)

-r, --rounds <n> – number of rounds to play (default: 1)

--help – show usage

📸 Example Output
text
🧠 Memory Trainer
Difficulty: medium (6 digits)
Round 1/3
Remember: 4 8 2 1 5 9
(Wait 2 seconds...)
Enter the numbers (space‑separated): 4 8 2 1 5 9
✅ Correct!
Score: 1/1
📁 Repository Structure
text
.
├── README.md
├── python/
│   └── memory_trainer.py
├── go/
│   └── memory_trainer.go
├── javascript/
│   └── memory_trainer.js
├── ruby/
│   └── memory_trainer.rb
├── php/
│   └── memory_trainer.php
├── java/
│   └── MemoryTrainer.java
├── csharp/
│   └── MemoryTrainer.cs
└── cpp/
    └── memory_trainer.cpp
