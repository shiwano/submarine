require 'rails_helper'

RSpec.describe RoomMember, type: :model do
  subject { create(:room_member) }

  it { should belong_to :user }
  it { should belong_to(:room).counter_cache(true) }
end
