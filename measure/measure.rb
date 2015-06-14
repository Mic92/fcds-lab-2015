#!/usr/bin/env ruby

require 'open3'

def bench(cmd)
  puts "$ #{cmd}"
  cpus = File.read("/proc/cpuinfo").scan(/processor/).size
  (1..cpus).each do |i|
    times = []
    10.times do |n|
      taskset = "taskset -c 1-#{i} bash -c '#{cmd}'"
      time = ""
      Open3.popen3(taskset) do |_, _, stderr, _|
        time = stderr.read
      end
      #puts time
      time =~ /([\d.]+)ms/
      times << $1.to_f
    end
    times.delete(times.min)
    times.delete(times.max)
    avg = times.reduce(:+) / times.size
    puts "#{i},#{avg}"
  end
end

bench "./fcds-lab-2015 bucketsort < lab/c_sequential/bucketsort/input/medium.in > out"
bench "./fcds-lab-2015 threesat < lab/c_sequential/3sat/input/large.in"
