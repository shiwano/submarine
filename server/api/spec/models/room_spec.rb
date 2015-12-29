require 'rails_helper'

RSpec.describe Room, type: :model do
  let(:room) { create(:room) }
  subject { room }

  it { should have_many(:room_members).dependent(:delete_all) }
  it { should have_many :users }

  it { should validate_presence_of :battle_server_base_uri }

  describe '.joinable' do
    let(:joinable_rooms) { [ create(:room), create(:room) ] }
    subject { Room.joinable }

    before do
      joinable_rooms
      create(:room, :full)
    end
    it 'should return joinable rooms' do
      expect(subject).to eq joinable_rooms
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
        expect { subject }.to raise_error ApplicationError::RoomAlreadyJoined
      end
    end

    context 'when the room is full' do
      before do
        Room.max_room_members_count.times do
          user = create(:user)
          room.join_user!(user)
        end
      end
      it 'should raise error' do
        expect { subject }.to raise_error ApplicationError::RoomIsFull
      end
    end
  end

  describe '#to_room_api_type' do
    subject { room.to_room_api_type }
    it { should be_a_kind_of TyphenApi::Model::Submarine::Room }
  end

  describe '#to_joined_room_api_type' do
    let(:room) { create(:room, :with_user) }
    subject { room.to_joined_room_api_type(room.users.first) }
    it { should be_a_kind_of TyphenApi::Model::Submarine::JoinedRoom }
  end
end
