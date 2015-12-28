require 'rails_helper'

RSpec.describe Room, type: :model do
  subject { create(:room) }

  it { should have_many(:room_members).dependent(:delete_all) }
  it { should have_many :users }

  describe '#join' do
    let(:user) { create(:user) }

    it 'should join user' do
      expect { subject.join(user) }.to change { subject.users.count }.from(0).to(1)
    end

    context 'when the room is full' do
      before do
        subject.max_member_count.times do
          user = create(:user)
          subject.join(user)
        end
      end

      it 'should raise the error' do
        expect { subject.join(user) }.to raise_error ApplicationError::RoomIsFull
      end
    end
  end

  describe '#to_api_type' do
    it 'should return an instance of the typhen api type' do
      expect(subject.to_api_type).to be_a_kind_of TyphenApi::Model::Submarine::Room
    end
  end
end
