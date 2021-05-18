class Anagram
  attr_reader :word

  def initialize(word)
    @word = word
  end

  def match(anagram_candidates)
    anagram_candidates.select do |candidate|
      sum_each_char(candidate) == word_char_to_count && is_different?(candidate)
    end
  end


  private

  def sum_each_char(word_)
    word_.downcase.chars.each_with_object(Hash.new(0)) do |char, memo|
      memo[char] += 1
      memo
    end
  end

  def word_char_to_count
    @word_char_to_count ||= sum_each_char(word)
  end

  def is_different?(other_word)
    word.downcase != other_word.downcase
  end

end
