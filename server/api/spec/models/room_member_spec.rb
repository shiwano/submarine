require 'rails_helper'

RSpec.describe RoomMember, type: :model do
  let(:room_member) { create(:room_member) }
  subject { room_member }

  it { should belong_to :user }
  it { should belong_to(:room).counter_cache(true) }

  it { should validate_presence_of :room_key }

  describe '#to_battle_room_member_api_type' do
    subject { room_member.to_battle_room_member_api_type }
    it { should be_a_kind_of TyphenApi::Model::Submarine::Battle::RoomMember }
  end
end
