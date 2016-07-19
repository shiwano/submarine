require 'rails_helper'

RSpec.describe AccessToken, type: :model do
  let(:access_token) { create(:access_token) }
  subject { access_token }

  it { should belong_to :user }

  it { should validate_presence_of(:user) }
  it { should validate_uniqueness_of(:user) }
  it { should validate_presence_of(:token) }
  it { should validate_uniqueness_of(:token) }
  it { should validate_presence_of(:expires_at) }

  describe '#generate_token' do
    subject { access_token.generate_token }

    it { is_expected.to match /\A([a-f0-9]{2}){64}\z/i }
    it 'should renew expires_at' do
      expect { subject }.to change { access_token.expires_at }
    end
  end

  describe '#expired?' do
    let(:expires_at) { Time.now }
    let(:access_token) { create(:access_token, expires_at: expires_at) }
    subject { access_token.expired? }
    after { Timecop.return }

    context 'when current time is over expires_at' do
      before { Timecop.freeze(expires_at + 1.second) }
      it { is_expected.to be true }
    end
    context 'when current time is not over expires_at' do
      before { Timecop.freeze(expires_at - 1.second) }
      it { is_expected.to be false }
    end
  end
end
