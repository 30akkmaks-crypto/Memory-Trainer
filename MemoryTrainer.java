// MemoryTrainer.java
import java.io.*;
import java.util.*;
import java.util.concurrent.*;

public class MemoryTrainer {
    private static final int SHOW_TIME = 2; // seconds

    public static void clearScreen() {
        System.out.print("\033[2J\033[0;0H");
        System.out.flush();
    }

    public static int[] generateSequence(int length) {
        int[] seq = new int[length];
        Random rand = new Random();
        for (int i=0; i<length; i++) seq[i] = rand.nextInt(10);
        return seq;
    }

    public static void displaySequence(int[] seq) {
        System.out.println("Remember these numbers:");
        for (int n : seq) System.out.print(n + " ");
        System.out.println();
        try {
            Thread.sleep(SHOW_TIME * 1000);
        } catch (InterruptedException e) {}
        clearScreen();
    }

    public static int[] getUserInput(int length) throws IOException {
        BufferedReader reader = new BufferedReader(new InputStreamReader(System.in));
        while (true) {
            System.out.print("Enter the " + length + " numbers (space‑separated): ");
            String line = reader.readLine();
            String[] parts = line.trim().split("\\s+");
            if (parts.length != length) {
                System.out.println("Please enter exactly " + length + " numbers.");
                continue;
            }
            int[] nums = new int[length];
            boolean ok = true;
            for (int i=0; i<length; i++) {
                try {
                    nums[i] = Integer.parseInt(parts[i]);
                } catch (NumberFormatException e) {
                    System.out.println("Please enter valid integers.");
                    ok = false;
                    break;
                }
            }
            if (ok) return nums;
        }
    }

    public static boolean playRound(int length) throws IOException {
        int[] seq = generateSequence(length);
        displaySequence(seq);
        int[] userSeq = getUserInput(length);
        return Arrays.equals(seq, userSeq);
    }

    public static void main(String[] args) throws Exception {
        Map<String, String> params = new HashMap<>();
        for (int i=0; i<args.length; i++) {
            if (args[i].startsWith("--")) {
                String key = args[i].substring(2);
                if (i+1 < args.length && !args[i+1].startsWith("--")) {
                    params.put(key, args[++i]);
                } else {
                    params.put(key, "");
                }
            }
        }
        int length = 0;
        String difficulty = params.get("difficulty");
        int rounds = Integer.parseInt(params.getOrDefault("rounds", "1"));

        if (params.containsKey("length")) {
            length = Integer.parseInt(params.get("length"));
        } else if (difficulty != null) {
            switch (difficulty) {
                case "easy": length = 4; break;
                case "medium": length = 6; break;
                case "hard": length = 8; break;
                default: System.out.println("Invalid difficulty."); return;
            }
        } else {
            length = 6;
        }
        if (length < 1 || rounds < 1) {
            System.out.println("Length and rounds must be at least 1.");
            return;
        }

        System.out.println("\n🧠 Memory Trainer");
        System.out.printf("Difficulty: %s (%d digits)%n", difficulty != null ? difficulty : "custom", length);
        System.out.println("Rounds: " + rounds);

        int correct = 0;
        for (int r=1; r<=rounds; r++) {
            System.out.printf("%nRound %d/%d%n", r, rounds);
            if (playRound(length)) {
                System.out.println("✅ Correct!");
                correct++;
            } else {
                System.out.println("❌ Wrong.");
            }
        }
        System.out.printf("%nScore: %d/%d%n", correct, rounds);
    }
}
