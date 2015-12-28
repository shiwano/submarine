# == Schema Information
#
# Table name: rooms
#
#  id                     :integer          not null, primary key
#  battle_server_base_uri :string(255)
#  lock_version           :integer
#  created_at             :datetime         not null
#  updated_at             :datetime         not null
#

class Room < ActiveRecord::Base
  has_many :room_members
  has_many :users, through: :room_members
end
