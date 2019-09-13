class Complement
  DNA_RNA_MAP = {
    'G' => 'C',
    'C' => 'G',
    'T' => 'A',
    'A' => 'U'
  }

  def self.of_dna(nucleodite)
    return '' if nucleodite.empty?
    return self::DNA_RNA_MAP[nucleodite] if nucleodite.length == 1

    output = ''
    nucleodite.each_char do |c|
      output = output + self::DNA_RNA_MAP[c]
    end
    output
  end
end
