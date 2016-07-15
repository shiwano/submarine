require 'rails_helper'

RSpec.describe FindUserController, type: :controller do
  context 'POST service', with_login: true do
    before do
      create(:user, name: 'Shimakaze')
      post :service, params: params
    end

    describe '#user' do
      subject { assigns(:user) }

      context 'with an existing user name' do
        let(:params) { { name: 'Shimakaze' } }
        it { should be_a_kind_of(User) }
        it { should have_attributes(persisted?: true, name: 'Shimakaze') }
      end

      context 'with a no-existing user name' do
        let(:params) { { name: 'Kaga' } }
        it { should be_nil }
      end
    end
  end
end
