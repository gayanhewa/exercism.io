class Acronym
  ABR = {
    'PNG' => 'Portable Network Graphics',
    'ROR' => 'Ruby on Rails',
    'FIFO' => 'First In, First Out',
    'GIMP' => 'GNU Image Manipulation Program',
    'CMOS' => 'Complementary metal-oxide semiconductor',
    'ROTFLSHTMDCOALM' => 'Rolling On The Floor Laughing So Hard That My Dogs Came Over And Licked Me',
    'SIMUFTA' => 'Something - I made up from thin air'
  }

  def self.abbreviate(str)
    self::ABR.key(str)
  end
end
