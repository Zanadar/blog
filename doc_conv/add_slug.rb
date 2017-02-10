#!/usr/bin/ruby
require 'pry'

Dir.glob('/Users/zanadar/code/zblog/content/post/*.md') do |file|
  File.
foreach(file) do |line|
    mat = /title = (.*)$/.match line
    binding.pry


  end


end
