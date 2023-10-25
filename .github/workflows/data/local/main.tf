resource "random_pet" "pet" {}

output "pet_name" {
  value = random_pet.pet.id
}
