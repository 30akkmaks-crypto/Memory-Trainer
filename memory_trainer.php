# memory_trainer.php
#!/usr/bin/env php
<?php

function clearScreen() {
    system('clear') || system('cls');
}

function generateSequence($length) {
    $seq = [];
    for ($i=0; $i<$length; $i++) {
        $seq[] = rand(0,9);
    }
    return $seq;
}

function displaySequence($seq, $showTime) {
    echo "Remember these numbers:\n";
    echo implode(' ', $seq) . "\n";
    sleep($showTime);
    clearScreen();
}

function getUserInput($length) {
    while (true) {
        echo "Enter the $length numbers (space‑separated): ";
        $line = trim(fgets(STDIN));
        $parts = preg_split('/\s+/', $line);
        if (count($parts) != $length) {
            echo "Please enter exactly $length numbers.\n";
            continue;
        }
        $nums = array_map('intval', $parts);
        if (count($nums) == $length) {
            return $nums;
        }
        echo "Please enter valid integers.\n";
    }
}

function playRound($length, $showTime) {
    $seq = generateSequence($length);
    displaySequence($seq, $showTime);
    $userSeq = getUserInput($length);
    return $seq == $userSeq;
}

$opts = getopt("l:d:r:", ["length:", "difficulty:", "rounds:"]);
$length = isset($opts['l']) ? (int)$opts['l'] : (isset($opts['length']) ? (int)$opts['length'] : 0);
$difficulty = $opts['d'] ?? $opts['difficulty'] ?? null;
$rounds = isset($opts['r']) ? (int)$opts['r'] : (isset($opts['rounds']) ? (int)$opts['rounds'] : 1);

if ($length > 0) {
    $seqLen = $length;
} elseif ($difficulty) {
    switch ($difficulty) {
        case 'easy': $seqLen = 4; break;
        case 'medium': $seqLen = 6; break;
        case 'hard': $seqLen = 8; break;
        default: echo "Invalid difficulty.\n"; exit(1);
    }
} else {
    $seqLen = 6;
}
if ($seqLen < 1 || $rounds < 1) {
    echo "Length and rounds must be at least 1.\n";
    exit(1);
}

echo "\n🧠 Memory Trainer\n";
echo "Difficulty: " . ($difficulty ?: 'custom') . " ($seqLen digits)\n";
echo "Rounds: $rounds\n";

$correct = 0;
$showTime = 2;
for ($r=1; $r<=$rounds; $r++) {
    echo "\nRound $r/$rounds\n";
    if (playRound($seqLen, $showTime)) {
        echo "✅ Correct!\n";
        $correct++;
    } else {
        echo "❌ Wrong.\n";
    }
}
echo "\nScore: $correct/$rounds\n";
?>
