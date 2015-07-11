require 'rails_helper'

RSpec.describe LoginController, type: :controller do

  context 'POST service' do
    let(:user) { create(:user, :with_stupid_password) }
    let(:params) { { name: user.name, password: 'secret' } }

    before do
      post :service, params
    end

    describe 'logged_in_user' do
      subject { assigns(:logged_in_user) }
      it { should be_a_kind_of(User) }
      it { should have_attributes(id: user.id, persisted?: true) }
    end
  end

end
