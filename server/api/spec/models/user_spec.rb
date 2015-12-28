require 'rails_helper'

RSpec.describe User, type: :model do
  let(:user) { create(:user) }
  subject { user }

  it { should have_one :room_member }
  it { should have_one :room }

  it { should validate_length_of(:password).is_at_least(6) }
  it { should validate_presence_of :name }
  it { should validate_uniqueness_of :name }
  it { should validate_length_of(:name).is_at_least(3) }

  describe '#create_room' do
    let(:params) { { battle_server_base_uri: Faker::Internet.ip_v4_address } }
    subject { user.create_room(params) }

    context 'when the user has no room' do
      it 'should create a room' do
        expect(subject).to be_a_kind_of Room
      end
      it 'should join into the created room' do
        expect(subject.users).to include(user)
      end
    end

    context 'when the user has a room' do
      before do
        user.create_room(params)
        user.reload
      end
      it 'should raise error' do
        expect { subject }.to raise_error ApplicationError::RoomCreatingFailed
      end
    end
  end

  describe '#to_api_type' do
    subject { user.to_api_type }
    it { should be_a_kind_of TyphenApi::Model::Submarine::User }
  end
end
