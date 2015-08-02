require 'rails_helper'

RSpec.describe Config, type: :model do
  subject { Config }

  it 'should delegate method calls to a TyphenApi::Model::Submarine::Config instance' do
    expect(subject.version).to be_a_kind_of String
  end
end
