# memory_trainer.rb
#!/usr/bin/env ruby
require 'optparse'

def clear_screen
  system('clear') || system('cls')
end

def generate_sequence(length)
  Array.new(length) { rand(10) }
end

def display_sequence(seq, show_time)
  puts "Remember these numbers:"
  puts seq.join(' ')
  sleep(show_time)
  clear_screen
end

def get_user_input(length)
  loop do
    print "Enter the #{length} numbers (space‑separated): "
    input = gets.chomp
    parts = input.split
    if parts.length != length
      puts "Please enter exactly #{length} numbers."
      next
    end
    begin
      nums = parts.map(&:to_i)
      return nums
    rescue
      puts "Please enter valid integers."
    end
  end
end

def play_round(length, show_time)
  seq = generate_sequence(length)
  display_sequence(seq, show_time)
  user_seq = get_user_input(length)
  seq == user_seq
end

options = {}
OptionParser.new do |opts|
  opts.banner = "Usage: memory_trainer.rb [options]"
  opts.on('-l LENGTH', '--length LENGTH', Integer, 'Number of digits per sequence') { |v| options[:length] = v }
  opts.on('-d DIFFICULTY', '--difficulty DIFFICULTY', ['easy', 'medium', 'hard'], 'Preset lengths') { |v| options[:difficulty] = v }
  opts.on('-r ROUNDS', '--rounds ROUNDS', Integer, 'Number of rounds') { |v| options[:rounds] = v }
end.parse!

seq_len = if options[:length]
            options[:length]
          elsif options[:difficulty]
            case options[:difficulty]
            when 'easy' then 4
            when 'medium' then 6
            when 'hard' then 8
            end
          else
            6
          end

rounds = options[:rounds] || 1

if seq_len < 1 || rounds < 1
  puts "Length and rounds must be at least 1."
  exit 1
end

puts "\n🧠 Memory Trainer"
puts "Difficulty: #{options[:difficulty] || 'custom'} (#{seq_len} digits)"
puts "Rounds: #{rounds}"

correct = 0
show_time = 2
rounds.times do |r|
  puts "\nRound #{r+1}/#{rounds}"
  if play_round(seq_len, show_time)
    puts "✅ Correct!"
    correct += 1
  else
    puts "❌ Wrong."
  end
end
puts "\nScore: #{correct}/#{rounds}"
