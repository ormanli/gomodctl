# This file was generated by GoReleaser. DO NOT EDIT.
class Gomodctl < Formula
  desc "search,check and update go modules"
  homepage "https://github.com/beatlabs/gomodctl"
  version "0.4.0"
  bottle :unneeded

  if OS.mac?
    url "https://github.com/beatlabs/gomodctl/releases/download/v0.4.0/gomodctl_Darwin_x86_64.tar.gz"
    sha256 "2a0bd638eae601f8be38ec02b6e80305b47e9a685b3f9f90f63c6d09e0a4e717"
  elsif OS.linux?
    if Hardware::CPU.intel?
      url "https://github.com/beatlabs/gomodctl/releases/download/v0.4.0/gomodctl_Linux_x86_64.tar.gz"
      sha256 "5cd09abc01e1e27e4d01f1ccc9fef81ecaed03a95b9ec1ae454dd3c940cdd87c"
    end
  end

  def install
    bin.install "gomodctl"
  end

  test do
    system "#{bin/gomodctl}"
  end
end
