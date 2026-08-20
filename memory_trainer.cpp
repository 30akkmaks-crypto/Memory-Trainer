// memory_trainer.cpp
#include <iostream>
#include <vector>
#include <string>
#include <sstream>
#include <cstdlib>
#include <ctime>
#include <thread>
#include <chrono>
#include <getopt.h>

#ifdef _WIN32
#include <windows.h>
#else
#include <unistd.h>
#endif

using namespace std;

void clearScreen() {
#ifdef _WIN32
    system("cls");
#else
    system("clear");
#endif
}

vector<int> generateSequence(int length) {
    vector<int> seq(length);
    for (int i=0; i<length; i++) seq[i] = rand() % 10;
    return seq;
}

void displaySequence(const vector<int>& seq, int showTime) {
    cout << "Remember these numbers:" << endl;
    for (int n : seq) cout << n << " ";
    cout << endl;
    this_thread::sleep_for(chrono::seconds(showTime));
    clearScreen();
}

vector<int> getUserInput(int length) {
    while (true) {
        cout << "Enter the " << length << " numbers (space‑separated): ";
        string line;
        getline(cin, line);
        stringstream ss(line);
        vector<int> nums;
        int num;
        while (ss >> num) nums.push_back(num);
        if ((int)nums.size() != length) {
            cout << "Please enter exactly " << length << " numbers." << endl;
            continue;
        }
        return nums;
    }
}

bool playRound(int length, int showTime) {
    vector<int> seq = generateSequence(length);
    displaySequence(seq, showTime);
    vector<int> userSeq = getUserInput(length);
    return seq == userSeq;
}

int main(int argc, char* argv[]) {
    static struct option long_options[] = {
        {"length", required_argument, 0, 'l'},
        {"difficulty", required_argument, 0, 'd'},
        {"rounds", required_argument, 0, 'r'},
        {0,0,0,0}
    };
    int opt;
    int length = 0;
    string difficulty;
    int rounds = 1;
    while ((opt = getopt_long(argc, argv, "l:d:r:", long_options, nullptr)) != -1) {
        switch (opt) {
            case 'l': length = stoi(optarg); break;
            case 'd': difficulty = optarg; break;
            case 'r': rounds = stoi(optarg); break;
            default:
                cerr << "Usage: memory_trainer --length <n> --difficulty <easy|medium|hard> --rounds <n>" << endl;
                return 1;
        }
    }

    if (length == 0) {
        if (difficulty == "easy") length = 4;
        else if (difficulty == "medium") length = 6;
        else if (difficulty == "hard") length = 8;
        else length = 6;
    }
    if (length < 1 || rounds < 1) {
        cerr << "Length and rounds must be at least 1." << endl;
        return 1;
    }

    srand(time(nullptr));
    cout << "\n🧠 Memory Trainer" << endl;
    cout << "Difficulty: " << (difficulty.empty() ? "custom" : difficulty) << " (" << length << " digits)" << endl;
    cout << "Rounds: " << rounds << endl;

    int correct = 0;
    int showTime = 2;
    for (int r=1; r<=rounds; r++) {
        cout << "\nRound " << r << "/" << rounds << endl;
        if (playRound(length, showTime)) {
            cout << "✅ Correct!" << endl;
            correct++;
        } else {
            cout << "❌ Wrong." << endl;
        }
    }
    cout << "\nScore: " << correct << "/" << rounds << endl;
    return 0;
}
