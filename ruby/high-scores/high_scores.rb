class HighScores
  def initialize(scores)
    @scores = scores
  end

  def scores
    @scores
  end

  def latest
    @scores.last
  end

  def personal_best
    @scores.max
  end

  def personal_top_three
    @scores.sort.reverse[0..2]
  end

  def latest_is_personal_best?
    self.latest == self.personal_best
  end
end
