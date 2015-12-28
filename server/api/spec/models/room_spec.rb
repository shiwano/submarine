require 'rails_helper'

RSpec.describe Room, type: :model do
  let(:room) { create(:room) }
  subject { room }

  it { should have_many(:room_members).dependent(:delete_all) }
  it { should have_many :users }

  describe '#join' do
    let(:user) { create(:user) }
    subject { room.join(user) }

    it 'should join an user' do
      expect { subject }.to change { room.users.count }.from(0).to(1)
    end

    context 'when the room is full' do
      before do
        room.max_member_count.times do
          user = create(:user)
          room.join(user)
        end
      end
      it 'should raise error' do
        expect { subject }.to raise_error ApplicationError::RoomIsFull
      end
    end
  end

  describe '#to_api_type' do
    subject { room.to_api_type }
    it { should be_a_kind_of TyphenApi::Model::Submarine::Room }
  end
end
