class ResistorColorDuo
  COLORS = ['black', 'brown', 'red', 'orange', 'yellow', 'green', 'blue', 'violet', 'grey', 'white']

  def self.value(colors=[])
    return nil if colors.empty?
    return self::COLORS.index(colors[0]) if colors.length == 1
    "#{self::COLORS.index(colors[0])}#{self::COLORS.index(colors[1])}".to_i
  end
end
