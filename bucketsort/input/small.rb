#!/usr/bin/env ruby

puts 94 * 2
('!'..'~').to_a.reverse.each do |c1|
  ('a'..'b').to_a.reverse.each do |c2|
    print c1 * 5, c2 * 2, "\n"
  end
end
