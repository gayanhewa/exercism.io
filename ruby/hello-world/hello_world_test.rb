class HelloWorldTest < Minitest::Test
  def test_say_hi
    # skip
    assert_equal 'Hello, World!', HelloWorld.new.hello
  end
end
