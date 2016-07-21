require 'rails_helper'

RSpec.describe Room, type: :model do
  let(:room) { create(:room) }
  subject { room }

  it { should have_many(:room_members).dependent(:delete_all) }
  it { should have_many :users }

  it { should validate_presence_of :battle_server_base_uri }

  describe '.joinable' do
    subject { Room.joinable }

    before do
      @joinable_rooms = [ create(:room), create(:room) ]
      create(:room, :full)
    end
    it 'should return joinable rooms' do
      expect(subject).to eq @joinable_rooms
    end
  end

  describe '#join_user!' do
    let(:user) { create(:user) }
    subject { room.join_user!(user) }

    it 'should join an user' do
      expect { subject }.to change { room.users.count }.from(0).to(1)
    end

    context 'with an user that has already joined into a room' do
      before do
        room.join_user!(user)
      end
      it 'should raise error' do
        expect { subject }.to raise_error GameError::RoomAlreadyJoined
      end
    end
    context 'when the room is full' do
      before do
        Room.max_room_members_count.times do
          room.join_user!(create(:user))
        end
      end
      it 'should raise error' do
        expect { subject }.to raise_error GameError::RoomIsFull
      end
    end
  end

  describe '#as_room_api_type' do
    subject { room.as_room_api_type }
    it { should be_a_kind_of TyphenApi::Model::Submarine::Room }
  end

  describe '#as_battle_room_api_type' do
    subject { room.as_battle_room_api_type }
    it { should be_a_kind_of TyphenApi::Model::Submarine::Battle::Room }
  end

  describe '#as_joined_room_api_type' do
    let(:room) { create(:room, :with_user) }
    subject { room.as_joined_room_api_type(room.users.first) }
    it { should be_a_kind_of TyphenApi::Model::Submarine::JoinedRoom }
  end
end
