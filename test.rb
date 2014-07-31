require 'minitest/autorun'

describe 'catd' do
  before do
    system "rm -rf ./tmp/*"
    system "mkdir -p tmp/foo.d"
    system "echo bar > tmp/foo.d/bar"
    system "echo baz > tmp/foo.d/baz"
  end

  it do
    system "go run main.go -once -dir tmp/foo.d -file tmp/foo"
    File.read("tmp/foo").must_equal "bar\nbaz\n"
  end
end
