shared_context 'with login', with_login: true do
  let(:current_user) { create(:user, :with_stupid_password) }

  before do
    login_user(current_user)
  end
end
