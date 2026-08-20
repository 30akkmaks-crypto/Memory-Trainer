// MemoryTrainer.cs
using System;
using System.Collections.Generic;
using System.Linq;

class MemoryTrainer
{
    const int ShowTime = 2;

    static void ClearScreen()
    {
        Console.Clear();
    }

    static int[] GenerateSequence(int length)
    {
        var rand = new Random();
        return Enumerable.Range(0, length).Select(_ => rand.Next(10)).ToArray();
    }

    static void DisplaySequence(int[] seq)
    {
        Console.WriteLine("Remember these numbers:");
        Console.WriteLine(string.Join(" ", seq));
        System.Threading.Thread.Sleep(ShowTime * 1000);
        ClearScreen();
    }

    static int[] GetUserInput(int length)
    {
        while (true)
        {
            Console.Write($"Enter the {length} numbers (space‑separated): ");
            var line = Console.ReadLine().Trim();
            var parts = line.Split(' ', StringSplitOptions.RemoveEmptyEntries);
            if (parts.Length != length)
            {
                Console.WriteLine($"Please enter exactly {length} numbers.");
                continue;
            }
            int[] nums = new int[length];
            bool ok = true;
            for (int i=0; i<length; i++)
            {
                if (!int.TryParse(parts[i], out nums[i]))
                {
                    Console.WriteLine("Please enter valid integers.");
                    ok = false;
                    break;
                }
            }
            if (ok) return nums;
        }
    }

    static bool PlayRound(int length)
    {
        var seq = GenerateSequence(length);
        DisplaySequence(seq);
        var userSeq = GetUserInput(length);
        return seq.SequenceEqual(userSeq);
    }

    static void Main(string[] args)
    {
        var parsed = ParseArgs(args);
        int length = 0;
        string difficulty = parsed.GetValueOrDefault("difficulty");
        int rounds = int.Parse(parsed.GetValueOrDefault("rounds", "1"));

        if (parsed.ContainsKey("length"))
        {
            length = int.Parse(parsed["length"]);
        }
        else if (difficulty != null)
        {
            switch (difficulty)
            {
                case "easy": length = 4; break;
                case "medium": length = 6; break;
                case "hard": length = 8; break;
                default: Console.WriteLine("Invalid difficulty."); return;
            }
        }
        else
        {
            length = 6;
        }
        if (length < 1 || rounds < 1)
        {
            Console.WriteLine("Length and rounds must be at least 1.");
            return;
        }

        Console.WriteLine("\n🧠 Memory Trainer");
        Console.WriteLine($"Difficulty: {(difficulty ?? "custom")} ({length} digits)");
        Console.WriteLine($"Rounds: {rounds}");

        int correct = 0;
        for (int r=1; r<=rounds; r++)
        {
            Console.WriteLine($"\nRound {r}/{rounds}");
            if (PlayRound(length))
            {
                Console.WriteLine("✅ Correct!");
                correct++;
            }
            else
            {
                Console.WriteLine("❌ Wrong.");
            }
        }
        Console.WriteLine($"\nScore: {correct}/{rounds}");
    }

    static Dictionary<string, string> ParseArgs(string[] args)
    {
        var dict = new Dictionary<string, string>();
        for (int i=0; i<args.Length; i++)
        {
            if (args[i].StartsWith("--"))
            {
                string key = args[i].Substring(2);
                if (i+1 < args.Length && !args[i+1].StartsWith("--"))
                    dict[key] = args[++i];
                else
                    dict[key] = "";
            }
        }
        return dict;
    }
}
