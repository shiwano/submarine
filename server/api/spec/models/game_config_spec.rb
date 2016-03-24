require 'rails_helper'

RSpec.describe GameConfig, type: :model do
  subject { GameConfig }

  its(:battle_server_base_uri) { is_expected.to be_a_kind_of String }
end
