// memory_trainer.js
#!/usr/bin/env node
const readline = require('readline');
const { program } = require('commander');

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

function clearScreen() {
    process.stdout.write('\x1b[2J\x1b[0;0H');
}

function generateSequence(length) {
    return Array.from({ length }, () => Math.floor(Math.random() * 10));
}

function displaySequence(seq, showTime) {
    console.log('Remember these numbers:');
    console.log(seq.join(' '));
    setTimeout(() => clearScreen(), showTime * 1000);
    return new Promise(resolve => setTimeout(resolve, showTime * 1000));
}

function getUserInput(length) {
    return new Promise((resolve) => {
        rl.question(`Enter the ${length} numbers (space‑separated): `, (answer) => {
            const parts = answer.trim().split(/\s+/);
            if (parts.length !== length) {
                console.log(`Please enter exactly ${length} numbers.`);
                resolve(getUserInput(length));
                return;
            }
            const nums = parts.map(Number);
            if (nums.some(isNaN)) {
                console.log('Please enter valid integers.');
                resolve(getUserInput(length));
                return;
            }
            resolve(nums);
        });
    });
}

async function playRound(length, showTime) {
    const seq = generateSequence(length);
    await displaySequence(seq, showTime);
    const userSeq = await getUserInput(length);
    return seq.every((v, i) => v === userSeq[i]);
}

async function main() {
    program
        .option('-l, --length <n>', 'Number of digits per sequence', parseInt)
        .option('-d, --difficulty <type>', 'easy, medium, hard')
        .option('-r, --rounds <n>', 'Number of rounds', parseInt, 1)
        .parse(process.argv);

    const opts = program.opts();
    let seqLen;
    if (opts.length) {
        seqLen = opts.length;
    } else if (opts.difficulty) {
        switch (opts.difficulty) {
            case 'easy': seqLen = 4; break;
            case 'medium': seqLen = 6; break;
            case 'hard': seqLen = 8; break;
            default: console.error('Invalid difficulty.'); process.exit(1);
        }
    } else {
        seqLen = 6;
    }
    if (seqLen < 1 || opts.rounds < 1) {
        console.error('Length and rounds must be at least 1.');
        process.exit(1);
    }

    console.log(`\n🧠 Memory Trainer`);
    console.log(`Difficulty: ${opts.difficulty || 'custom'} (${seqLen} digits)`);
    console.log(`Rounds: ${opts.rounds}`);

    let correct = 0;
    const showTime = 2;
    for (let r = 1; r <= opts.rounds; r++) {
        console.log(`\nRound ${r}/${opts.rounds}`);
        const ok = await playRound(seqLen, showTime);
        if (ok) {
            console.log('✅ Correct!');
            correct++;
        } else {
            console.log('❌ Wrong.');
        }
    }
    console.log(`\nScore: ${correct}/${opts.rounds}`);
    rl.close();
}

main().catch(console.error);
