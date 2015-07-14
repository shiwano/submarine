require 'rails_helper'

RSpec.describe SignUpController, type: :controller do
  context 'POST service' do
    let(:params) { { name: Faker::Name.first_name, password: Faker::Internet.password(6, 20) } }

    before do
      post :service, params
    end

    describe 'new_user' do
      subject { @controller.new_user }
      it { should be_a_kind_of(User) }
      it { should have_attributes(persisted?: true) }
    end
  end
end
