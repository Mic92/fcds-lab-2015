#!/usr/bin/env ruby

require 'open3'

def delete_first(ary, elem)
  idx = ary.index(elem)
  return if idx.nil?
  ary.delete_at(idx)
end

def delete_biggest_smallest(ary)
  delete_first(ary, ary.min)
  delete_first(ary, ary.max)
end

def bench(cmd)
  puts "$ #{cmd}"
  cpus = File.read("/proc/cpuinfo").scan(/processor/).size
  (1..cpus).each do |i|
    compute_times = []
    overall_times = []
    10.times do |n|
      taskset = "taskset -c 1-#{i} bash -c '#{cmd}'"
      time = ""
      Open3.popen3(taskset) do |_, _, stderr, _|
        time = stderr.read
      end
      $stderr.puts("#{n}: #{time}")
      time =~ /computation time: ([\d.]+)ms, overall time: ([\d.]+)ms/
      compute_times << $1.to_f
      overall_times << $2.to_f
      system("sync")
    end

    delete_biggest_smallest(compute_times)
    delete_biggest_smallest(overall_times)

    compute_avg = compute_times.reduce(:+) / compute_times.size
    overall_avg = overall_times.reduce(:+) / overall_times.size
    puts "#{i},#{compute_avg},#{overall_avg}"
  end
end

Dir.chdir(File.realpath(File.join(File.dirname(__FILE__), "..")))

unless File.exists?("lab/c_sequential/haar/input/large.in")
  $stderr.puts "Please run input_generator for haar first like this:\n" +
    "  $ cd lab/c_sequential/haar/ && make && bin/input_generator input/large.in"
  exit(1)
end

bench "./fcds-lab-2015 bucketsort < lab/c_sequential/bucketsort/input/medium.in > out"
bench "./fcds-lab-2015 threesat < lab/c_sequential/3sat/input/large.in"
bench "./fcds-lab-2015 haar < lab/c_sequential/haar/input/large.in > haar.out"
