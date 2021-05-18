class Series
  attr_reader :slice_string

  def initialize(slice_string)
    @slice_string = slice_string
  end
  
  def slices(n)
    raise ArgumentError if n > string_length
    
    loop_count_by(n: n).times.map do |i|
      slice_string[i, n]
    end
  end


  private

  def loop_count_by(n:)
    string_length - (n-1)
  end

  def string_length
    @string_length ||= slice_string.length
  end

end
