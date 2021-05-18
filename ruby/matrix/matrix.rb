class Matrix
  attr_reader :rows

  def initialize(str)
    @rows = build_array_from_string(str)
  end

  def columns
    rows.transpose
  end


  private
  
  def build_array_from_string(str)
    str.each_line.map(&:strip).map { |s| s.split.map(&:to_i) }
  end

end
