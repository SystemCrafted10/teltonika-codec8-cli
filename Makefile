# Define the password variable
pw ?=

# Define the hashed password 
HASHED_PASSWORD = a6f6b908d98bb8de5d31318569f976f35c58a1b8e44cf3dacd9c72e061a9b83b10a7d9988026563c18f8f837dd573ecaa95a74086ea583aad092852a0572e224


# Define the paths to the scripts
ENCRYPT_SCRIPT = ./scripts/encrypt.sh
DECRYPT_SCRIPT = ./scripts/decrypt.sh

# Define paths to docker-compose files
DOCKER_COMPOSE_FILE = ./dockerfiles/track.iotsolutions.shared/docker-compose.yml

# Check if password is provided
ifeq ($(pw),)
$(error PASSWORD is not set. Usage: make target pw=<password>)
endif

# Function to hash the password
define CHECK_PASSWORD
  echo -n "$(1)" | sha512sum | awk '{print $$1}'
endef

# Check if the provided password matches the hashed password
CHECKED_PASSWORD := $(shell $(call CHECK_PASSWORD,$(pw)))
ifeq ($(CHECKED_PASSWORD),$(HASHED_PASSWORD))
# Password is correct, continue
else
$(error PASSWORD does not match)
endif

# --------------------------------------------------
# Encryption and Decryption Targets
# --------------------------------------------------

# Target to encrypt .env files and notes/*.md files
encrypt:
	chmod +x $(ENCRYPT_SCRIPT)
	$(ENCRYPT_SCRIPT) $(pw)

# Target to decrypt .env files and notes/*.md files
decrypt:
	chmod +x $(DECRYPT_SCRIPT)
	$(DECRYPT_SCRIPT) $(pw)

# --------------------------------------------------
# Docker Compose Targets
# --------------------------------------------------

# Target to bring up docker services
docker-up:
	docker-compose -f $(DOCKER_COMPOSE_FILE) up -d

# Target to bring down docker services
docker-down:
	docker-compose -f $(DOCKER_COMPOSE_FILE) down

# --------------------------------------------------
# Default target (if needed)
# --------------------------------------------------

all: encrypt decrypt

# 	------------------------------------------------

#	make encrypt pw=<password>

#	make decrypt pw=<password>

# 	------------------------------------------------

