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

  describe '.no_expired' do
    subject { AccessToken.no_expired }

    before do
      @no_expired_access_tokens = [create(:access_token), create(:access_token)]
      create(:access_token, :expired)
    end
    it 'should return no-expired access tokens' do
      expect(subject).to eq @no_expired_access_tokens
    end
  end

  describe '.find_by_token' do
    subject { AccessToken.find_by_token(token) }

    context 'with a valid token' do
      let(:token) { access_token.token }
      it { is_expected.to eq access_token }
    end
    context 'with a invalid token' do
      let(:token) { 'invalid' }
      it { is_expected.to be nil }
    end
    context 'when the access token is expired' do
      let(:access_token) { create(:access_token, :expired) }
      let(:token) { access_token.token }
      it { is_expected.to be nil }
    end
  end

  describe '#generate_token' do
    subject { access_token.generate_token }

    it { is_expected.to match /\A([a-f0-9]{2}){32}\z/i }
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
