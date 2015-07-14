require 'rails_helper'

RSpec.describe User, type: :model do
  subject { create(:user) }

  it { should validate_length_of(:password).is_at_least(6) }

  it { should validate_presence_of :name }
  it { should validate_uniqueness_of :name }
  it { should validate_length_of(:name).is_at_least(3) }

  describe '#to_api_type' do
    it 'should return an instance of TyphenApi::Model::Submarine::User' do
      expect(subject.to_api_type).to be_a_kind_of TyphenApi::Model::Submarine::User
    end
  end
end
