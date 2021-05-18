class Transpose
  
  def self.transpose(str)
    str.split("\n").map { |e| e.split('') }.transpose.map { |e| e.join }.join("\n")
  end
end
