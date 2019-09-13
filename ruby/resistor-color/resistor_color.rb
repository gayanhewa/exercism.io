class ResistorColor
  COLORS = ['black', 'brown', 'red', 'orange', 'yellow', 'green', 'blue', 'violet', 'grey', 'white']

  def self.color_code(color)
    return self::COLORS.index(color)
  end
end
