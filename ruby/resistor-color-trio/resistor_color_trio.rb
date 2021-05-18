class ResistorColorTrio
  attr_reader :colors

  COLOR_TO_VALUE = {
    'black' => 0,
    'brown' => 1,
    'red' => 2,
    'orange' => 3,
    'yellow' => 4,
    'green' => 5,
    'blue' => 6,
    'violet' => 7,
    'grey' => 8,
    'white' => 9
  }

  COLOR_TO_MULTIPLIER = {
    'black' => 1,
    'brown' => 10,
    'red' => 100,
    'orange' => 1_000,
    'yellow' => 10_000,
    'green' => 100_000,
    'blue' => 1_000_000,
    'violet' => 10_000_000,
    'grey' => 100_000_0000,
    'white' => 1_000_000_000
  }

  def initialize(colors)
    raise ArgumentError.new unless colors.all? { |color| COLOR_TO_VALUE.keys.include? color }
    @colors = colors
  end

  def label
  "Resistor value: #{humanized}ohms"
  end


  private

  def humanized
    res = value / 1_000
    res >= 1 ? "#{res} kilo" : "#{value} " 
  end

  def value
    "#{first_color_value}#{second_color_value}".to_i * third_color_value
  end

  def first_color_value
    COLOR_TO_VALUE[colors.first]
  end

  def second_color_value
    COLOR_TO_VALUE[colors[1]]
  end

  def third_color_value
    COLOR_TO_MULTIPLIER[colors[2]]
  end

end
