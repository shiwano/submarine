require 'rails_helper'

RSpec.describe Config, type: :model do
  subject { Config }

  its(:battle_server_base_uri) { is_expected.to be_a_kind_of String }
end
