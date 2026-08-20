# memory_trainer.py
import sys
import random
import time
import argparse
import os

def clear_screen():
    os.system('cls' if os.name == 'nt' else 'clear')

def generate_sequence(length):
    return [random.randint(0, 9) for _ in range(length)]

def display_sequence(seq, show_time=2):
    print("Remember these numbers:")
    print(' '.join(map(str, seq)))
    time.sleep(show_time)
    clear_screen()

def get_user_input(length):
    while True:
        try:
            inp = input(f"Enter the {length} numbers (space‑separated): ").strip().split()
            if len(inp) != length:
                print(f"Please enter exactly {length} numbers.")
                continue
            return list(map(int, inp))
        except ValueError:
            print("Please enter valid integers.")

def play_round(length, show_time=2):
    seq = generate_sequence(length)
    display_sequence(seq, show_time)
    user_seq = get_user_input(length)
    return user_seq == seq

def main():
    parser = argparse.ArgumentParser(description="Memory Trainer - Numbers")
    parser.add_argument('-l', '--length', type=int, help='Number of digits per sequence')
    parser.add_argument('-d', '--difficulty', choices=['easy', 'medium', 'hard'],
                        help='Preset lengths: easy=4, medium=6, hard=8')
    parser.add_argument('-r', '--rounds', type=int, default=1, help='Number of rounds')
    args = parser.parse_args()

    # Determine length
    if args.length:
        length = args.length
    elif args.difficulty:
        length = {'easy': 4, 'medium': 6, 'hard': 8}[args.difficulty]
    else:
        length = 6  # default medium

    if length < 1:
        print("Length must be at least 1.")
        sys.exit(1)
    if args.rounds < 1:
        print("Rounds must be at least 1.")
        sys.exit(1)

    print(f"\n🧠 Memory Trainer")
    print(f"Difficulty: {args.difficulty or 'custom'} ({length} digits)")
    print(f"Rounds: {args.rounds}")

    correct = 0
    for r in range(1, args.rounds + 1):
        print(f"\nRound {r}/{args.rounds}")
        if play_round(length):
            print("✅ Correct!")
            correct += 1
        else:
            print("❌ Wrong.")
    print(f"\nScore: {correct}/{args.rounds}")

if __name__ == "__main__":
    main()
