class AddOwnerToSites < ActiveRecord::Migration[8.0]
  def change
    add_reference :sites, :owner, polymorphic: true, null: false, index: true
    add_reference :sites, :creator, foreign_key: { to_table: :users }, null: false, index: true
  end
end
