class Gigasecond
  def self.from(now=Time.now)
    now + 1000000000 if !now.nil?
  end
end
