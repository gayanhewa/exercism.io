class Complement
  def self.of_dna(nucleodite)
    return '' if nucleodite.empty?
    nucleodite.tr('GCTA', 'CGAU')
  end
end
