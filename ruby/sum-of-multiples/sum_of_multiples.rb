class SumOfMultiples
  attr_reader :multiples

  def initialize(*multiples)
    @multiples = multiples
  end

  def to(upper_limit)
    1.upto(upper_limit.pred).select do |number|
      multiples.any? { |multiple| number % multiple == 0 }
    end.sum
  end

end
