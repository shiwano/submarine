require 'rails_helper'

RSpec.describe Room, type: :model do
  subject { create(:room) }

  it { should have_many :room_members }
  it { should have_many :users }
end
