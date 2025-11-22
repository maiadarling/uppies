class ApplicationController < ActionController::API
  def current_user
    # Just gonna stub first user for now until I figure out how I want auth to work
    User.first
  end
end
