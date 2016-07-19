require 'rails_helper'

RSpec.describe User, type: :model do
  let(:user) { create(:user) }
  subject { user }

  it { should have_one :access_token }
  it { should have_one :room_member }
  it { should have_one :room }

  it { should validate_presence_of(:encrypted_auth_token) }
  it { should validate_uniqueness_of(:encrypted_auth_token) }
  it { should validate_presence_of(:name) }
  it { should validate_length_of(:name).is_at_most(16) }

  describe '.salt' do
    subject { User.salt }
    it { is_expected.to eq Rails.application.secrets[:secret_key_base] }
  end

  describe '.encrypt_auth_token' do
    let(:expected_encrypted_auth_token) {
      '457b54850539ab90751a6ae79c76874eef7102f310338694509bd605d601bb6a' +
      'bd0a17f1efb48c1bed7c3675970b3c3493ca6aa2357cfa7e5dc82e77b624b9d5'
    }
    subject { User.encrypt_auth_token('test') }
    it { is_expected.to eq expected_encrypted_auth_token }
  end

  describe '.find_by_auth_token' do
    subject { User.find_by_auth_token(auth_token) }
    before do
      @user = create(:user, :with_stupid_auth_token)
    end

    context 'with the valid access token' do
      let(:auth_token) { 'secret' }
      it { is_expected.to eq @user }
    end
    context 'with the invalid access token' do
      let(:auth_token) { 'invalid' }
      it { is_expected.to be nil }
    end
  end

  describe '.find_by_access_token' do
    subject { User.find_by_access_token(access_token) }

    context 'with the valid access token' do
      let(:access_token) { user.generate_access_token! }
      it { is_expected.to eq user }
    end
    context 'with the invalid access token' do
      let(:access_token) { 'invalid' }
      it { is_expected.to be nil }
    end
  end

  describe '#generate_auth_token' do
    subject { user.generate_auth_token }
    it { is_expected.to match /\A([a-f0-9]{2}){64}\z/i }
    it 'should set the generated auth token to auth_token' do
      expect(user).to receive(:auth_token=).and_call_original
      subject
    end
  end

  describe '#generate_access_token!' do
    subject { user.generate_access_token! }
    it { is_expected.to match /\A([a-f0-9]{2}){64}\z/i }
    it 'should save the access token' do
      expect { subject }.to change {
        user.access_token.try(:persisted?) || false
      }.from(false).to(true)
    end
  end

  describe '#auth_token=' do
    before do
      allow(user.class).to receive(:encrypt_auth_token).and_return('encrypted')
    end
    it 'should set the given auth token to encrypted_auth_token' do
      user.auth_token = 'test'
      expect(user.encrypted_auth_token).to eq 'encrypted'
    end
  end

  describe '#create_room!' do
    subject { user.create_room! }

    context 'when the user has no room' do
      it 'should create a room' do
        expect(subject).to be_a_kind_of Room
      end
      it 'should join into the created room' do
        expect(subject.users).to include(user)
      end
    end

    context 'when the user has a room' do
      let(:user) { create(:user, :with_room) }

      it 'should raise error' do
        expect { subject }.to raise_error GameError::RoomAlreadyJoined
      end
    end
  end

  describe '#as_user_api_type' do
    subject { user.as_user_api_type }
    it { should be_a_kind_of TyphenApi::Model::Submarine::User }
  end

  describe '#as_logged_in_user_api_type' do
    subject { user.as_logged_in_user_api_type }
    it { should be_a_kind_of TyphenApi::Model::Submarine::LoggedInUser }
  end
end
