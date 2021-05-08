/*
   GoToSocial
   Copyright (C) 2021 GoToSocial Authors admin@gotosocial.org

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package testrig

import (
	"crypto/rand"
	"crypto/rsa"
	"net"
	"time"

	"github.com/superseriousbusiness/gotosocial/internal/db/gtsmodel"
	"github.com/superseriousbusiness/gotosocial/internal/oauth"
)

// NewTestTokens returns a map of tokens keyed according to which account the token belongs to.
func NewTestTokens() map[string]*oauth.Token {
	tokens := map[string]*oauth.Token{
		"local_account_1": {
			ID:              "64cf4214-33ab-4220-b5ca-4a6a12263b20",
			ClientID:        "73b48d42-029d-4487-80fc-329a5cf67869",
			UserID:          "44e36b79-44a4-4bd8-91e9-097f477fe97b",
			RedirectURI:     "http://localhost:8080",
			Scope:           "read write follow push",
			Access:          "NZAZOTC0OWITMDU0NC0ZODG4LWE4NJITMWUXM2M4MTRHZDEX",
			AccessCreateAt:  time.Now(),
			AccessExpiresAt: time.Now().Add(72 * time.Hour),
		},
		"local_account_2": {
			ID:              "b04cae99-39b5-4610-a425-dc6b91c78a72",
			ClientID:        "a4f6a2ea-a32b-4600-8853-72fc4ad98a1f",
			UserID:          "d120bd97-866f-4a05-9690-a1294b9934c3",
			RedirectURI:     "http://localhost:8080",
			Scope:           "read write follow push",
			Access:          "PIPINALKNNNFNF98717NAMNAMNFKIJKJ881818KJKJAKJJJA",
			AccessCreateAt:  time.Now(),
			AccessExpiresAt: time.Now().Add(72 * time.Hour),
		},
	}
	return tokens
}

// NewTestClients returns a map of Clients keyed according to which account they are used by.
func NewTestClients() map[string]*oauth.Client {
	clients := map[string]*oauth.Client{
		"admin_account": {
			ID:     "1c5cefc8-f0c9-4307-8506-ca6e3888675e",
			Secret: "dda8e835-2c9c-4bd2-9b8b-77c2e26d7a7a",
			Domain: "http://localhost:8080",
			UserID: "0fb02eae-2214-473f-9667-0a43f22d75ff", // admin_account
		},
		"local_account_1": {
			ID:     "73b48d42-029d-4487-80fc-329a5cf67869",
			Secret: "c3724c74-dc3b-41b2-a108-0ea3d8399830",
			Domain: "http://localhost:8080",
			UserID: "44e36b79-44a4-4bd8-91e9-097f477fe97b", // local_account_1
		},
		"local_account_2": {
			ID:     "a4f6a2ea-a32b-4600-8853-72fc4ad98a1f",
			Secret: "8f5603a5-c721-46cd-8f1b-2e368f51379f",
			Domain: "http://localhost:8080",
			UserID: "d120bd97-866f-4a05-9690-a1294b9934c3", // local_account_2
		},
	}
	return clients
}

// NewTestApplications returns a map of applications keyed to which number application they are.
func NewTestApplications() map[string]*gtsmodel.Application {
	apps := map[string]*gtsmodel.Application{
		"admin_account": {
			ID:           "9bf9e368-037f-444d-8ffd-1091d1c21c4c",
			Name:         "superseriousbusiness",
			Website:      "https://superserious.business",
			RedirectURI:  "http://localhost:8080",
			ClientID:     "1c5cefc8-f0c9-4307-8506-ca6e3888675e", // admin client
			ClientSecret: "dda8e835-2c9c-4bd2-9b8b-77c2e26d7a7a", // admin client
			Scopes:       "read write follow push",
			VapidKey:     "76ae0095-8a10-438f-9f49-522d1985b190",
		},
		"application_1": {
			ID:           "f88697b8-ee3d-46c2-ac3f-dbb85566c3cc",
			Name:         "really cool gts application",
			Website:      "https://reallycool.app",
			RedirectURI:  "http://localhost:8080",
			ClientID:     "73b48d42-029d-4487-80fc-329a5cf67869", // client_1
			ClientSecret: "c3724c74-dc3b-41b2-a108-0ea3d8399830", // client_1
			Scopes:       "read write follow push",
			VapidKey:     "4738dfd7-ca73-4aa6-9aa9-80e946b7db36",
		},
		"application_2": {
			ID:           "6b0cd164-8497-4cd5-bec9-957886fac5df",
			Name:         "kindaweird",
			Website:      "https://kindaweird.app",
			RedirectURI:  "http://localhost:8080",
			ClientID:     "a4f6a2ea-a32b-4600-8853-72fc4ad98a1f", // client_2
			ClientSecret: "8f5603a5-c721-46cd-8f1b-2e368f51379f", // client_2
			Scopes:       "read write follow push",
			VapidKey:     "c040a5fc-e1e2-4859-bbea-0a3efbca1c4b",
		},
	}
	return apps
}

// NewTestUsers returns a map of Users keyed by which account belongs to them.
func NewTestUsers() map[string]*gtsmodel.User {
	users := map[string]*gtsmodel.User{
		"unconfirmed_account": {
			ID:                     "0f7b1d24-1e49-4ee0-bc7e-fd87b7289eea",
			Email:                  "",
			AccountID:              "59e197f5-87cd-4be8-ac7c-09082ccc4b4d",
			EncryptedPassword:      "$2y$10$ggWz5QWwnx6kzb9g0tnIJurFtE0dhr5Zfeaqs9iFuUIXzafQlJVZS", // 'password'
			CreatedAt:              time.Now(),
			SignUpIP:               net.ParseIP("199.222.111.89"),
			UpdatedAt:              time.Time{},
			CurrentSignInAt:        time.Time{},
			CurrentSignInIP:        nil,
			LastSignInAt:           time.Time{},
			LastSignInIP:           nil,
			SignInCount:            0,
			InviteID:               "",
			ChosenLanguages:        []string{},
			FilteredLanguages:      []string{},
			Locale:                 "en",
			CreatedByApplicationID: "",
			LastEmailedAt:          time.Time{},
			ConfirmationToken:      "a5a280bd-34be-44a3-8330-a57eaf61b8dd",
			ConfirmedAt:            time.Time{},
			ConfirmationSentAt:     time.Now(),
			UnconfirmedEmail:       "weed_lord420@example.org",
			Moderator:              false,
			Admin:                  false,
			Disabled:               false,
			Approved:               false,
			ResetPasswordToken:     "",
			ResetPasswordSentAt:    time.Time{},
		},
		"admin_account": {
			ID:                     "0fb02eae-2214-473f-9667-0a43f22d75ff",
			Email:                  "admin@example.org",
			AccountID:              "8020dbb4-1e7b-4d99-a872-4cf94e64210f",
			EncryptedPassword:      "$2y$10$ggWz5QWwnx6kzb9g0tnIJurFtE0dhr5Zfeaqs9iFuUIXzafQlJVZS", // 'password'
			CreatedAt:              time.Now().Add(-72 * time.Hour),
			SignUpIP:               net.ParseIP("89.22.189.19"),
			UpdatedAt:              time.Now().Add(-72 * time.Hour),
			CurrentSignInAt:        time.Now().Add(-10 * time.Minute),
			CurrentSignInIP:        net.ParseIP("89.122.255.1"),
			LastSignInAt:           time.Now().Add(-2 * time.Hour),
			LastSignInIP:           net.ParseIP("89.122.255.1"),
			SignInCount:            78,
			InviteID:               "",
			ChosenLanguages:        []string{"en"},
			FilteredLanguages:      []string{},
			Locale:                 "en",
			CreatedByApplicationID: "",
			LastEmailedAt:          time.Now().Add(-30 * time.Minute),
			ConfirmationToken:      "",
			ConfirmedAt:            time.Now().Add(-72 * time.Hour),
			ConfirmationSentAt:     time.Time{},
			UnconfirmedEmail:       "",
			Moderator:              true,
			Admin:                  true,
			Disabled:               false,
			Approved:               true,
			ResetPasswordToken:     "",
			ResetPasswordSentAt:    time.Time{},
		},
		"local_account_1": {
			ID:                     "44e36b79-44a4-4bd8-91e9-097f477fe97b",
			Email:                  "zork@example.org",
			AccountID:              "580072df-4d03-4684-a412-89fd6f7d77e6",
			EncryptedPassword:      "$2y$10$ggWz5QWwnx6kzb9g0tnIJurFtE0dhr5Zfeaqs9iFuUIXzafQlJVZS", // 'password'
			CreatedAt:              time.Now().Add(-36 * time.Hour),
			SignUpIP:               net.ParseIP("59.99.19.172"),
			UpdatedAt:              time.Now().Add(-72 * time.Hour),
			CurrentSignInAt:        time.Now().Add(-30 * time.Minute),
			CurrentSignInIP:        net.ParseIP("88.234.118.16"),
			LastSignInAt:           time.Now().Add(-2 * time.Hour),
			LastSignInIP:           net.ParseIP("147.111.231.154"),
			SignInCount:            9,
			InviteID:               "",
			ChosenLanguages:        []string{"en"},
			FilteredLanguages:      []string{},
			Locale:                 "en",
			CreatedByApplicationID: "f88697b8-ee3d-46c2-ac3f-dbb85566c3cc",
			LastEmailedAt:          time.Now().Add(-55 * time.Minute),
			ConfirmationToken:      "",
			ConfirmedAt:            time.Now().Add(-34 * time.Hour),
			ConfirmationSentAt:     time.Now().Add(-36 * time.Hour),
			UnconfirmedEmail:       "",
			Moderator:              false,
			Admin:                  false,
			Disabled:               false,
			Approved:               true,
			ResetPasswordToken:     "",
			ResetPasswordSentAt:    time.Time{},
		},
		"local_account_2": {
			ID:                     "f8d6272e-2d71-4d0c-97d3-2ba7a0b75bf7",
			Email:                  "tortle.dude@example.org",
			AccountID:              "eecaad73-5703-426d-9312-276641daa31e",
			EncryptedPassword:      "$2y$10$ggWz5QWwnx6kzb9g0tnIJurFtE0dhr5Zfeaqs9iFuUIXzafQlJVZS", // 'password'
			CreatedAt:              time.Now().Add(-36 * time.Hour),
			SignUpIP:               net.ParseIP("59.99.19.172"),
			UpdatedAt:              time.Now().Add(-72 * time.Hour),
			CurrentSignInAt:        time.Now().Add(-30 * time.Minute),
			CurrentSignInIP:        net.ParseIP("118.44.18.196"),
			LastSignInAt:           time.Now().Add(-2 * time.Hour),
			LastSignInIP:           net.ParseIP("198.98.21.15"),
			SignInCount:            9,
			InviteID:               "",
			ChosenLanguages:        []string{"en"},
			FilteredLanguages:      []string{},
			Locale:                 "en",
			CreatedByApplicationID: "",
			LastEmailedAt:          time.Now().Add(-55 * time.Minute),
			ConfirmationToken:      "",
			ConfirmedAt:            time.Now().Add(-34 * time.Hour),
			ConfirmationSentAt:     time.Now().Add(-36 * time.Hour),
			UnconfirmedEmail:       "",
			Moderator:              false,
			Admin:                  false,
			Disabled:               false,
			Approved:               true,
			ResetPasswordToken:     "",
			ResetPasswordSentAt:    time.Time{},
		},
	}

	return users
}

// NewTestAccounts returns a map of accounts keyed by what type of account they are.
func NewTestAccounts() map[string]*gtsmodel.Account {
	accounts := map[string]*gtsmodel.Account{
		"instance_account": {
			ID:       "39b745a3-774d-4b65-8bb2-b63d9e20a343",
			Username: "localhost:8080",
		},
		"unconfirmed_account": {
			ID:                      "59e197f5-87cd-4be8-ac7c-09082ccc4b4d",
			Username:                "weed_lord420",
			AvatarMediaAttachmentID: "",
			HeaderMediaAttachmentID: "",
			DisplayName:             "",
			Fields:                  []gtsmodel.Field{},
			Note:                    "",
			Memorial:                false,
			MovedToAccountID:        "",
			CreatedAt:               time.Now(),
			UpdatedAt:               time.Now(),
			Bot:                     false,
			Reason:                  "hi, please let me in! I'm looking for somewhere neato bombeato to hang out.",
			Locked:                  false,
			Discoverable:            false,
			Privacy:                 gtsmodel.VisibilityPublic,
			Sensitive:               false,
			Language:                "en",
			URI:                     "http://localhost:8080/users/weed_lord420",
			URL:                     "http://localhost:8080/@weed_lord420",
			LastWebfingeredAt:       time.Time{},
			InboxURL:                "http://localhost:8080/users/weed_lord420/inbox",
			OutboxURL:               "http://localhost:8080/users/weed_lord420/outbox",
			SharedInboxURL:          "",
			FollowersURL:            "http://localhost:8080/users/weed_lord420/followers",
			FeaturedCollectionURL:   "http://localhost:8080/users/weed_lord420/collections/featured",
			ActorType:               gtsmodel.ActivityStreamsPerson,
			AlsoKnownAs:             "",
			PrivateKey:              &rsa.PrivateKey{},
			PublicKey:               &rsa.PublicKey{},
			SensitizedAt:            time.Time{},
			SilencedAt:              time.Time{},
			SuspendedAt:             time.Time{},
			HideCollections:         false,
			SuspensionOrigin:        "",
		},
		"admin_account": {
			ID:                      "8020dbb4-1e7b-4d99-a872-4cf94e64210f",
			Username:                "admin",
			AvatarMediaAttachmentID: "",
			HeaderMediaAttachmentID: "",
			DisplayName:             "",
			Fields:                  []gtsmodel.Field{},
			Note:                    "",
			Memorial:                false,
			MovedToAccountID:        "",
			CreatedAt:               time.Now().Add(-72 * time.Hour),
			UpdatedAt:               time.Now().Add(-72 * time.Hour),
			Bot:                     false,
			Reason:                  "",
			Locked:                  false,
			Discoverable:            true,
			Privacy:                 gtsmodel.VisibilityPublic,
			Sensitive:               false,
			Language:                "en",
			URI:                     "http://localhost:8080/users/admin",
			URL:                     "http://localhost:8080/@admin",
			LastWebfingeredAt:       time.Time{},
			InboxURL:                "http://localhost:8080/users/admin/inbox",
			OutboxURL:               "http://localhost:8080/users/admin/outbox",
			SharedInboxURL:          "",
			FollowersURL:            "http://localhost:8080/users/admin/followers",
			FeaturedCollectionURL:   "http://localhost:8080/users/admin/collections/featured",
			ActorType:               gtsmodel.ActivityStreamsPerson,
			AlsoKnownAs:             "",
			PrivateKey:              &rsa.PrivateKey{},
			PublicKey:               &rsa.PublicKey{},
			SensitizedAt:            time.Time{},
			SilencedAt:              time.Time{},
			SuspendedAt:             time.Time{},
			HideCollections:         false,
			SuspensionOrigin:        "",
		},
		"local_account_1": {
			ID:                      "580072df-4d03-4684-a412-89fd6f7d77e6",
			Username:                "the_mighty_zork",
			AvatarMediaAttachmentID: "a849906f-8b8e-4b43-ac2f-6979ccbcd442",
			HeaderMediaAttachmentID: "",
			DisplayName:             "original zork (he/they)",
			Fields:                  []gtsmodel.Field{},
			Note:                    "hey yo this is my profile!",
			Memorial:                false,
			MovedToAccountID:        "",
			CreatedAt:               time.Now().Add(-48 * time.Hour),
			UpdatedAt:               time.Now().Add(-48 * time.Hour),
			Bot:                     false,
			Reason:                  "I wanna be on this damned webbed site so bad! Please! Wow",
			Locked:                  false,
			Discoverable:            true,
			Privacy:                 gtsmodel.VisibilityPublic,
			Sensitive:               false,
			Language:                "en",
			URI:                     "http://localhost:8080/users/the_mighty_zork",
			URL:                     "http://localhost:8080/@the_mighty_zork",
			LastWebfingeredAt:       time.Time{},
			InboxURL:                "http://localhost:8080/users/the_mighty_zork/inbox",
			OutboxURL:               "http://localhost:8080/users/the_mighty_zork/outbox",
			SharedInboxURL:          "",
			FollowersURL:            "http://localhost:8080/users/the_mighty_zork/followers",
			FeaturedCollectionURL:   "http://localhost:8080/users/the_mighty_zork/collections/featured",
			ActorType:               gtsmodel.ActivityStreamsPerson,
			AlsoKnownAs:             "",
			PrivateKey:              &rsa.PrivateKey{},
			PublicKey:               &rsa.PublicKey{},
			SensitizedAt:            time.Time{},
			SilencedAt:              time.Time{},
			SuspendedAt:             time.Time{},
			HideCollections:         false,
			SuspensionOrigin:        "",
		},
		"local_account_2": {
			ID:                      "eecaad73-5703-426d-9312-276641daa31e",
			Username:                "1happyturtle",
			AvatarMediaAttachmentID: "",
			HeaderMediaAttachmentID: "",
			DisplayName:             "happy little turtle :3",
			Fields:                  []gtsmodel.Field{},
			Note:                    "i post about things that concern me",
			Memorial:                false,
			MovedToAccountID:        "",
			CreatedAt:               time.Now().Add(-190 * time.Hour),
			UpdatedAt:               time.Now().Add(-36 * time.Hour),
			Bot:                     false,
			Reason:                  "",
			Locked:                  true,
			Discoverable:            false,
			Privacy:                 gtsmodel.VisibilityFollowersOnly,
			Sensitive:               false,
			Language:                "en",
			URI:                     "http://localhost:8080/users/1happyturtle",
			URL:                     "http://localhost:8080/@1happyturtle",
			LastWebfingeredAt:       time.Time{},
			InboxURL:                "http://localhost:8080/users/1happyturtle/inbox",
			OutboxURL:               "http://localhost:8080/users/1happyturtle/outbox",
			SharedInboxURL:          "",
			FollowersURL:            "http://localhost:8080/users/1happyturtle/followers",
			FeaturedCollectionURL:   "http://localhost:8080/users/1happyturtle/collections/featured",
			ActorType:               gtsmodel.ActivityStreamsPerson,
			AlsoKnownAs:             "",
			PrivateKey:              &rsa.PrivateKey{},
			PublicKey:               &rsa.PublicKey{},
			SensitizedAt:            time.Time{},
			SilencedAt:              time.Time{},
			SuspendedAt:             time.Time{},
			HideCollections:         false,
			SuspensionOrigin:        "",
		},
		"remote_account_1": {
			ID:       "c2c6e647-e2a9-4286-883b-e4a188186664",
			Username: "foss_satan",
			Domain:   "fossbros-anonymous.io",
			// AvatarFileName:        "http://localhost:8080/fileserver/media/eecaad73-5703-426d-9312-276641daa31e/avatar/original/d5e7c265-91a6-4d84-8c27-7e1efe5720da.jpeg",
			// AvatarContentType:     "image/jpeg",
			// AvatarFileSize:        0,
			// AvatarUpdatedAt:       time.Time{},
			// AvatarRemoteURL:       "",
			// HeaderFileName:        "http://localhost:8080/fileserver/media/eecaad73-5703-426d-9312-276641daa31e/header/original/e75d4117-21b6-4315-a428-eb3944235996.jpeg",
			// HeaderContentType:     "image/jpeg",
			// HeaderFileSize:        0,
			// HeaderUpdatedAt:       time.Time{},
			// HeaderRemoteURL:       "",
			DisplayName:           "big gerald",
			Fields:                []gtsmodel.Field{},
			Note:                  "i post about like, i dunno, stuff, or whatever!!!!",
			Memorial:              false,
			MovedToAccountID:      "",
			CreatedAt:             time.Now().Add(-190 * time.Hour),
			UpdatedAt:             time.Now().Add(-36 * time.Hour),
			Bot:                   false,
			Locked:                false,
			Discoverable:          true,
			Sensitive:             false,
			Language:              "en",
			URI:                   "https://fossbros-anonymous.io/users/foss_satan",
			URL:                   "https://fossbros-anonymous.io/@foss_satan",
			LastWebfingeredAt:     time.Time{},
			InboxURL:              "https://fossbros-anonymous.io/users/foss_satan/inbox",
			OutboxURL:             "https://fossbros-anonymous.io/users/foss_satan/outbox",
			SharedInboxURL:        "",
			FollowersURL:          "https://fossbros-anonymous.io/users/foss_satan/followers",
			FeaturedCollectionURL: "https://fossbros-anonymous.io/users/foss_satan/collections/featured",
			ActorType:             gtsmodel.ActivityStreamsPerson,
			AlsoKnownAs:           "",
			PrivateKey:            &rsa.PrivateKey{},
			PublicKey:             nil,
			SensitizedAt:          time.Time{},
			SilencedAt:            time.Time{},
			SuspendedAt:           time.Time{},
			HideCollections:       false,
			SuspensionOrigin:      "",
		},
		// "remote_account_2": {
		// 	ID:       "93287988-76c4-460f-9e68-a45b578bb6b2",
		// 	Username: "dailycatpics",
		// 	Domain:   "uwu.social",
		// },
		// "suspended_local_account": {
		// 	ID:       "e8a5cf4e-4b10-45a4-ad82-b6e37a09100d",
		// 	Username: "jeffbadman",
		// },
		// "suspended_remote_account": {
		// 	ID:       "17e6e09e-855d-4bf8-a1c3-7e780269f215",
		// 	Username: "ipfreely",
		// 	Domain:   "a-very-bad-website.com",
		// },
	}

	// generate keys for each account
	for _, v := range accounts {
		priv, err := rsa.GenerateKey(rand.Reader, 2048)
		if err != nil {
			panic(err)
		}
		pub := &priv.PublicKey

		// only local accounts get a private key
		if v.Domain == "" {
			v.PrivateKey = priv
		}
		v.PublicKey = pub
	}
	return accounts
}

// NewTestAttachments returns a map of attachments keyed according to which account
// and status they belong to, and which attachment number of that status they are.
func NewTestAttachments() map[string]*gtsmodel.MediaAttachment {
	return map[string]*gtsmodel.MediaAttachment{
		"admin_account_status_1_attachment_1": {
			ID:        "b052241b-f30f-4dc6-92fc-2bad0be1f8d8",
			StatusID:  "502ccd6f-0edf-48d7-9016-2dfa4d3714cd",
			URL:       "http://localhost:8080/fileserver/8020dbb4-1e7b-4d99-a872-4cf94e64210f/attachment/original/b052241b-f30f-4dc6-92fc-2bad0be1f8d8.jpeg",
			RemoteURL: "",
			CreatedAt: time.Now().Add(-71 * time.Hour),
			UpdatedAt: time.Now().Add(-71 * time.Hour),
			Type:      gtsmodel.FileTypeImage,
			FileMeta: gtsmodel.FileMeta{
				Original: gtsmodel.Original{
					Width:  1200,
					Height: 630,
					Size:   756000,
					Aspect: 1.9047619047619047,
				},
				Small: gtsmodel.Small{
					Width:  256,
					Height: 134,
					Size:   34304,
					Aspect: 1.9104477611940298,
				},
			},
			AccountID:         "8020dbb4-1e7b-4d99-a872-4cf94e64210f",
			Description:       "Black and white image of some 50's style text saying: Welcome On Board",
			ScheduledStatusID: "",
			Blurhash:          "LNJRdVM{00Rj%Mayt7j[4nWBofRj",
			Processing:        2,
			File: gtsmodel.File{
				Path:        "/gotosocial/storage/8020dbb4-1e7b-4d99-a872-4cf94e64210f/attachment/original/b052241b-f30f-4dc6-92fc-2bad0be1f8d8.jpeg",
				ContentType: "image/jpeg",
				FileSize:    62529,
				UpdatedAt:   time.Now().Add(-71 * time.Hour),
			},
			Thumbnail: gtsmodel.Thumbnail{
				Path:        "/gotosocial/storage/8020dbb4-1e7b-4d99-a872-4cf94e64210f/attachment/small/b052241b-f30f-4dc6-92fc-2bad0be1f8d8.jpeg",
				ContentType: "image/jpeg",
				FileSize:    6872,
				UpdatedAt:   time.Now().Add(-71 * time.Hour),
				URL:         "http://localhost:8080/fileserver/8020dbb4-1e7b-4d99-a872-4cf94e64210f/attachment/small/b052241b-f30f-4dc6-92fc-2bad0be1f8d8.jpeg",
				RemoteURL:   "",
			},
			Avatar: false,
			Header: false,
		},
		"local_account_1_status_4_attachment_1": {
			ID:        "510f6033-798b-4390-81b1-c38ca2205ad3",
			StatusID:  "18524c05-97dc-46d7-b474-c811bd9e1e32",
			URL:       "http://localhost:8080/fileserver/580072df-4d03-4684-a412-89fd6f7d77e6/attachment/original/510f6033-798b-4390-81b1-c38ca2205ad3.gif",
			RemoteURL: "",
			CreatedAt: time.Now().Add(-1 * time.Hour),
			UpdatedAt: time.Now().Add(-1 * time.Hour),
			Type:      gtsmodel.FileTypeGif,
			FileMeta: gtsmodel.FileMeta{
				Original: gtsmodel.Original{
					Width:  400,
					Height: 280,
					Size:   756000,
					Aspect: 1.4285714285714286,
				},
				Small: gtsmodel.Small{
					Width:  256,
					Height: 179,
					Size:   45824,
					Aspect: 1.4301675977653632,
				},
				Focus: gtsmodel.Focus{
					X: 0,
					Y: 0,
				},
			},
			AccountID:         "580072df-4d03-4684-a412-89fd6f7d77e6",
			Description:       "90's Trent Reznor turning to the camera",
			ScheduledStatusID: "",
			Blurhash:          "LEDara58O=t5EMSOENEN9]}?aK%0",
			Processing:        2,
			File: gtsmodel.File{
				Path:        "/gotosocial/storage/580072df-4d03-4684-a412-89fd6f7d77e6/attachment/original/510f6033-798b-4390-81b1-c38ca2205ad3.gif",
				ContentType: "image/gif",
				FileSize:    1109138,
				UpdatedAt:   time.Now().Add(-1 * time.Hour),
			},
			Thumbnail: gtsmodel.Thumbnail{
				Path:        "/gotosocial/storage/580072df-4d03-4684-a412-89fd6f7d77e6/attachment/small/510f6033-798b-4390-81b1-c38ca2205ad3.jpeg",
				ContentType: "image/jpeg",
				FileSize:    8803,
				UpdatedAt:   time.Now().Add(-1 * time.Hour),
				URL:         "http://localhost:8080/fileserver/580072df-4d03-4684-a412-89fd6f7d77e6/attachment/small/510f6033-798b-4390-81b1-c38ca2205ad3.jpeg",
				RemoteURL:   "",
			},
			Avatar: false,
			Header: false,
		},
		"local_account_1_unattached_1": {
			ID:        "7a3b9f77-ab30-461e-bdd8-e64bd1db3008",
			StatusID:  "", // this attachment isn't connected to a status YET
			URL:       "http://localhost:8080/fileserver/580072df-4d03-4684-a412-89fd6f7d77e6/attachment/original/7a3b9f77-ab30-461e-bdd8-e64bd1db3008.jpeg",
			RemoteURL: "",
			CreatedAt: time.Now().Add(30 * time.Second),
			UpdatedAt: time.Now().Add(30 * time.Second),
			Type:      gtsmodel.FileTypeGif,
			FileMeta: gtsmodel.FileMeta{
				Original: gtsmodel.Original{
					Width:  800,
					Height: 450,
					Size:   360000,
					Aspect: 1.7777777777777777,
				},
				Small: gtsmodel.Small{
					Width:  256,
					Height: 144,
					Size:   36864,
					Aspect: 1.7777777777777777,
				},
				Focus: gtsmodel.Focus{
					X: 0,
					Y: 0,
				},
			},
			AccountID:         "580072df-4d03-4684-a412-89fd6f7d77e6",
			Description:       "the oh you meme",
			ScheduledStatusID: "",
			Blurhash:          "LSAd]9ogDge-R:M|j=xWIto0xXWX",
			Processing:        2,
			File: gtsmodel.File{
				Path:        "/gotosocial/storage/580072df-4d03-4684-a412-89fd6f7d77e6/attachment/original/7a3b9f77-ab30-461e-bdd8-e64bd1db3008.jpeg",
				ContentType: "image/jpeg",
				FileSize:    27759,
				UpdatedAt:   time.Now().Add(30 * time.Second),
			},
			Thumbnail: gtsmodel.Thumbnail{
				Path:        "/gotosocial/storage/580072df-4d03-4684-a412-89fd6f7d77e6/attachment/small/7a3b9f77-ab30-461e-bdd8-e64bd1db3008.jpeg",
				ContentType: "image/jpeg",
				FileSize:    6177,
				UpdatedAt:   time.Now().Add(30 * time.Second),
				URL:         "http://localhost:8080/fileserver/580072df-4d03-4684-a412-89fd6f7d77e6/attachment/small/7a3b9f77-ab30-461e-bdd8-e64bd1db3008.jpeg",
				RemoteURL:   "",
			},
			Avatar: false,
			Header: false,
		},
		"local_account_1_avatar": {
			ID:        "a849906f-8b8e-4b43-ac2f-6979ccbcd442",
			StatusID:  "", // this attachment isn't connected to a status
			URL:       "http://localhost:8080/fileserver/580072df-4d03-4684-a412-89fd6f7d77e6/avatar/original/a849906f-8b8e-4b43-ac2f-6979ccbcd442.jpeg",
			RemoteURL: "",
			CreatedAt: time.Now().Add(47 * time.Hour),
			UpdatedAt: time.Now().Add(47 * time.Hour),
			Type:      gtsmodel.FileTypeImage,
			FileMeta: gtsmodel.FileMeta{
				Original: gtsmodel.Original{
					Width:  1092,
					Height: 1800,
					Size:   1965600,
					Aspect: 0.6066666666666667,
				},
				Small: gtsmodel.Small{
					Width:  155,
					Height: 256,
					Size:   39680,
					Aspect: 0.60546875,
				},
				Focus: gtsmodel.Focus{
					X: 0,
					Y: 0,
				},
			},
			AccountID:         "580072df-4d03-4684-a412-89fd6f7d77e6",
			Description:       "a green goblin looking nasty",
			ScheduledStatusID: "",
			Blurhash:          "LKK9MT,p|YSNDkJ-5rsmvnwcOoe:",
			Processing:        2,
			File: gtsmodel.File{
				Path:        "/gotosocial/storage/580072df-4d03-4684-a412-89fd6f7d77e6/avatar/original/a849906f-8b8e-4b43-ac2f-6979ccbcd442.jpeg",
				ContentType: "image/jpeg",
				FileSize:    457680,
				UpdatedAt:   time.Now().Add(47 * time.Hour),
			},
			Thumbnail: gtsmodel.Thumbnail{
				Path:        "/gotosocial/storage/580072df-4d03-4684-a412-89fd6f7d77e6/avatar/small/a849906f-8b8e-4b43-ac2f-6979ccbcd442.jpeg",
				ContentType: "image/jpeg",
				FileSize:    15374,
				UpdatedAt:   time.Now().Add(47 * time.Hour),
				URL:         "http://localhost:8080/fileserver/580072df-4d03-4684-a412-89fd6f7d77e6/avatar/small/a849906f-8b8e-4b43-ac2f-6979ccbcd442.jpeg",
				RemoteURL:   "",
			},
			Avatar: true,
			Header: false,
		},
	}
}

// NewTestEmojis returns a map of gts emojis, keyed by the emoji shortcode
func NewTestEmojis() map[string]*gtsmodel.Emoji {
	return map[string]*gtsmodel.Emoji{
		"rainbow": {
			ID:                   "a96ec4f3-1cae-47e4-a508-f9d66a6b221b",
			Shortcode:            "rainbow",
			Domain:               "",
			CreatedAt:            time.Now(),
			UpdatedAt:            time.Now(),
			ImageRemoteURL:       "",
			ImageStaticRemoteURL: "",
			ImageURL:             "http://localhost:8080/fileserver/39b745a3-774d-4b65-8bb2-b63d9e20a343/emoji/original/a96ec4f3-1cae-47e4-a508-f9d66a6b221b.png",
			ImagePath:            "/tmp/gotosocial/39b745a3-774d-4b65-8bb2-b63d9e20a343/emoji/original/a96ec4f3-1cae-47e4-a508-f9d66a6b221b.png",
			ImageStaticURL:       "http://localhost:8080/fileserver/39b745a3-774d-4b65-8bb2-b63d9e20a343/emoji/static/a96ec4f3-1cae-47e4-a508-f9d66a6b221b.png",
			ImageStaticPath:      "/tmp/gotosocial/39b745a3-774d-4b65-8bb2-b63d9e20a343/emoji/static/a96ec4f3-1cae-47e4-a508-f9d66a6b221b.png",
			ImageContentType:     "image/png",
			ImageFileSize:        36702,
			ImageStaticFileSize:  10413,
			ImageUpdatedAt:       time.Now(),
			Disabled:             false,
			URI:                  "http://localhost:8080/emoji/a96ec4f3-1cae-47e4-a508-f9d66a6b221b",
			VisibleInPicker:      true,
			CategoryID:           "",
		},
	}
}

type filenames struct {
	Original string
	Small    string
	Static   string
}

// newTestStoredAttachments returns a map of filenames, keyed according to which attachment they pertain to.
func newTestStoredAttachments() map[string]filenames {
	return map[string]filenames{
		"admin_account_status_1_attachment_1": {
			Original: "welcome-original.jpeg",
			Small:    "welcome-small.jpeg",
		},
		"local_account_1_status_4_attachment_1": {
			Original: "trent-original.gif",
			Small:    "trent-small.jpeg",
		},
		"local_account_1_unattached_1": {
			Original: "ohyou-original.jpeg",
			Small:    "ohyou-small.jpeg",
		},
		"local_account_1_avatar": {
			Original: "zork-original.jpeg",
			Small:    "zork-small.jpeg",
		},
	}
}

// newTestStoredEmoji returns a map of filenames, keyed according to which emoji they pertain to
func newTestStoredEmoji() map[string]filenames {
	return map[string]filenames{
		"rainbow": {
			Original: "rainbow-original.png",
			Static:   "rainbow-static.png",
		},
	}
}

// NewTestStatuses returns a map of statuses keyed according to which account
// and status they are.
func NewTestStatuses() map[string]*gtsmodel.Status {
	return map[string]*gtsmodel.Status{
		"admin_account_status_1": {
			ID:                       "502ccd6f-0edf-48d7-9016-2dfa4d3714cd",
			URI:                      "http://localhost:8080/users/admin/statuses/502ccd6f-0edf-48d7-9016-2dfa4d3714cd",
			URL:                      "http://localhost:8080/@admin/statuses/502ccd6f-0edf-48d7-9016-2dfa4d3714cd",
			Content:                  "hello world! #welcome ! first post on the instance :rainbow: !",
			Attachments:              []string{"b052241b-f30f-4dc6-92fc-2bad0be1f8d8"},
			Tags:                     []string{"a7e8f5ca-88a1-4652-8079-a187eab8d56e"},
			Mentions:                 []string{},
			Emojis:                   []string{"a96ec4f3-1cae-47e4-a508-f9d66a6b221b"},
			CreatedAt:                time.Now().Add(-71 * time.Hour),
			UpdatedAt:                time.Now().Add(-71 * time.Hour),
			Local:                    true,
			AccountID:                "8020dbb4-1e7b-4d99-a872-4cf94e64210f",
			InReplyToID:              "",
			BoostOfID:                "",
			ContentWarning:           "",
			Visibility:               gtsmodel.VisibilityPublic,
			Sensitive:                false,
			Language:                 "en",
			CreatedWithApplicationID: "9bf9e368-037f-444d-8ffd-1091d1c21c4c",
			VisibilityAdvanced: &gtsmodel.VisibilityAdvanced{
				Federated: true,
				Boostable: true,
				Replyable: true,
				Likeable:  true,
			},
			ActivityStreamsType: gtsmodel.ActivityStreamsNote,
		},
		"admin_account_status_2": {
			ID:                       "0fb3f1ac-5cd8-48ac-9050-3d95dc7e44e9",
			URI:                      "http://localhost:8080/users/admin/statuses/0fb3f1ac-5cd8-48ac-9050-3d95dc7e44e9",
			URL:                      "http://localhost:8080/@admin/statuses/0fb3f1ac-5cd8-48ac-9050-3d95dc7e44e9",
			Content:                  "🐕🐕🐕🐕🐕",
			CreatedAt:                time.Now().Add(-70 * time.Hour),
			UpdatedAt:                time.Now().Add(-70 * time.Hour),
			Local:                    true,
			AccountID:                "8020dbb4-1e7b-4d99-a872-4cf94e64210f",
			InReplyToID:              "",
			BoostOfID:                "",
			ContentWarning:           "open to see some puppies",
			Visibility:               gtsmodel.VisibilityPublic,
			Sensitive:                true,
			Language:                 "en",
			CreatedWithApplicationID: "9bf9e368-037f-444d-8ffd-1091d1c21c4c",
			VisibilityAdvanced: &gtsmodel.VisibilityAdvanced{
				Federated: true,
				Boostable: true,
				Replyable: true,
				Likeable:  true,
			},
			ActivityStreamsType: gtsmodel.ActivityStreamsNote,
		},
		"local_account_1_status_1": {
			ID:                       "91b1e795-74ff-4672-a4c4-476616710e2d",
			URI:                      "http://localhost:8080/users/the_mighty_zork/statuses/91b1e795-74ff-4672-a4c4-476616710e2d",
			URL:                      "http://localhost:8080/@the_mighty_zork/statuses/91b1e795-74ff-4672-a4c4-476616710e2d",
			Content:                  "hello everyone!",
			CreatedAt:                time.Now().Add(-47 * time.Hour),
			UpdatedAt:                time.Now().Add(-47 * time.Hour),
			Local:                    true,
			AccountID:                "580072df-4d03-4684-a412-89fd6f7d77e6",
			InReplyToID:              "",
			BoostOfID:                "",
			ContentWarning:           "introduction post",
			Visibility:               gtsmodel.VisibilityPublic,
			Sensitive:                true,
			Language:                 "en",
			CreatedWithApplicationID: "f88697b8-ee3d-46c2-ac3f-dbb85566c3cc",
			VisibilityAdvanced: &gtsmodel.VisibilityAdvanced{
				Federated: true,
				Boostable: true,
				Replyable: true,
				Likeable:  true,
			},
			ActivityStreamsType: gtsmodel.ActivityStreamsNote,
		},
		"local_account_1_status_2": {
			ID:                       "3dd328d9-8bb1-48f5-bc96-5ccc1c696b4c",
			URI:                      "http://localhost:8080/users/the_mighty_zork/statuses/3dd328d9-8bb1-48f5-bc96-5ccc1c696b4c",
			URL:                      "http://localhost:8080/@the_mighty_zork/statuses/3dd328d9-8bb1-48f5-bc96-5ccc1c696b4c",
			Content:                  "this is an unlocked local-only post that shouldn't federate, but it's still boostable, replyable, and likeable",
			CreatedAt:                time.Now().Add(-46 * time.Hour),
			UpdatedAt:                time.Now().Add(-46 * time.Hour),
			Local:                    true,
			AccountID:                "580072df-4d03-4684-a412-89fd6f7d77e6",
			InReplyToID:              "",
			BoostOfID:                "",
			ContentWarning:           "",
			Visibility:               gtsmodel.VisibilityUnlocked,
			Sensitive:                false,
			Language:                 "en",
			CreatedWithApplicationID: "f88697b8-ee3d-46c2-ac3f-dbb85566c3cc",
			VisibilityAdvanced: &gtsmodel.VisibilityAdvanced{
				Federated: false,
				Boostable: true,
				Replyable: true,
				Likeable:  true,
			},
			ActivityStreamsType: gtsmodel.ActivityStreamsNote,
		},
		"local_account_1_status_3": {
			ID:                       "5e41963f-8ab9-4147-9f00-52d56e19da65",
			URI:                      "http://localhost:8080/users/the_mighty_zork/statuses/5e41963f-8ab9-4147-9f00-52d56e19da65",
			URL:                      "http://localhost:8080/@the_mighty_zork/statuses/5e41963f-8ab9-4147-9f00-52d56e19da65",
			Content:                  "this is a very personal post that I don't want anyone to interact with at all, and i only want mutuals to see it",
			CreatedAt:                time.Now().Add(-45 * time.Hour),
			UpdatedAt:                time.Now().Add(-45 * time.Hour),
			Local:                    true,
			AccountID:                "580072df-4d03-4684-a412-89fd6f7d77e6",
			InReplyToID:              "",
			BoostOfID:                "",
			ContentWarning:           "test: you shouldn't be able to interact with this post in any way",
			Visibility:               gtsmodel.VisibilityMutualsOnly,
			Sensitive:                false,
			Language:                 "en",
			CreatedWithApplicationID: "f88697b8-ee3d-46c2-ac3f-dbb85566c3cc",
			VisibilityAdvanced: &gtsmodel.VisibilityAdvanced{
				Federated: true,
				Boostable: false,
				Replyable: false,
				Likeable:  false,
			},
			ActivityStreamsType: gtsmodel.ActivityStreamsNote,
		},
		"local_account_1_status_4": {
			ID:                       "18524c05-97dc-46d7-b474-c811bd9e1e32",
			URI:                      "http://localhost:8080/users/the_mighty_zork/statuses/18524c05-97dc-46d7-b474-c811bd9e1e32",
			URL:                      "http://localhost:8080/@the_mighty_zork/statuses/18524c05-97dc-46d7-b474-c811bd9e1e32",
			Content:                  "here's a little gif of trent",
			Attachments:              []string{"510f6033-798b-4390-81b1-c38ca2205ad3"},
			CreatedAt:                time.Now().Add(-1 * time.Hour),
			UpdatedAt:                time.Now().Add(-1 * time.Hour),
			Local:                    true,
			AccountID:                "580072df-4d03-4684-a412-89fd6f7d77e6",
			InReplyToID:              "",
			BoostOfID:                "",
			ContentWarning:           "eye contact, trent reznor gif",
			Visibility:               gtsmodel.VisibilityMutualsOnly,
			Sensitive:                false,
			Language:                 "en",
			CreatedWithApplicationID: "f88697b8-ee3d-46c2-ac3f-dbb85566c3cc",
			VisibilityAdvanced: &gtsmodel.VisibilityAdvanced{
				Federated: true,
				Boostable: true,
				Replyable: true,
				Likeable:  true,
			},
			ActivityStreamsType: gtsmodel.ActivityStreamsNote,
		},
		"local_account_2_status_1": {
			ID:                       "8945ccf2-3873-45e9-aa13-fd7163f19775",
			URI:                      "http://localhost:8080/users/1happyturtle/statuses/8945ccf2-3873-45e9-aa13-fd7163f19775",
			URL:                      "http://localhost:8080/@1happyturtle/statuses/8945ccf2-3873-45e9-aa13-fd7163f19775",
			Content:                  "🐢 hi everyone i post about turtles 🐢",
			CreatedAt:                time.Now().Add(-189 * time.Hour),
			UpdatedAt:                time.Now().Add(-189 * time.Hour),
			Local:                    true,
			AccountID:                "eecaad73-5703-426d-9312-276641daa31e",
			InReplyToID:              "",
			BoostOfID:                "",
			ContentWarning:           "introduction post",
			Visibility:               gtsmodel.VisibilityPublic,
			Sensitive:                true,
			Language:                 "en",
			CreatedWithApplicationID: "6b0cd164-8497-4cd5-bec9-957886fac5df",
			VisibilityAdvanced: &gtsmodel.VisibilityAdvanced{
				Federated: true,
				Boostable: true,
				Replyable: true,
				Likeable:  true,
			},
			ActivityStreamsType: gtsmodel.ActivityStreamsNote,
		},
		"local_account_2_status_2": {
			ID:                       "c7e25a86-f0d3-4705-a73c-c597f687d3dd",
			URI:                      "http://localhost:8080/users/1happyturtle/statuses/c7e25a86-f0d3-4705-a73c-c597f687d3dd",
			URL:                      "http://localhost:8080/@1happyturtle/statuses/c7e25a86-f0d3-4705-a73c-c597f687d3dd",
			Content:                  "🐢 this one is federated, likeable, and boostable but not replyable 🐢",
			CreatedAt:                time.Now().Add(-1 * time.Minute),
			UpdatedAt:                time.Now().Add(-1 * time.Minute),
			Local:                    true,
			AccountID:                "eecaad73-5703-426d-9312-276641daa31e",
			InReplyToID:              "",
			BoostOfID:                "",
			ContentWarning:           "",
			Visibility:               gtsmodel.VisibilityPublic,
			Sensitive:                true,
			Language:                 "en",
			CreatedWithApplicationID: "6b0cd164-8497-4cd5-bec9-957886fac5df",
			VisibilityAdvanced: &gtsmodel.VisibilityAdvanced{
				Federated: true,
				Boostable: true,
				Replyable: false,
				Likeable:  true,
			},
			ActivityStreamsType: gtsmodel.ActivityStreamsNote,
		},
		"local_account_2_status_3": {
			ID:                       "75960e30-7a8e-4f45-87fa-440a4d1c9572",
			URI:                      "http://localhost:8080/users/1happyturtle/statuses/75960e30-7a8e-4f45-87fa-440a4d1c9572",
			URL:                      "http://localhost:8080/@1happyturtle/statuses/75960e30-7a8e-4f45-87fa-440a4d1c9572",
			Content:                  "🐢 i don't mind people sharing this one but I don't want likes or replies to it because cba🐢",
			CreatedAt:                time.Now().Add(-2 * time.Minute),
			UpdatedAt:                time.Now().Add(-2 * time.Minute),
			Local:                    true,
			AccountID:                "eecaad73-5703-426d-9312-276641daa31e",
			InReplyToID:              "",
			BoostOfID:                "",
			ContentWarning:           "you won't be able to like or reply to this",
			Visibility:               gtsmodel.VisibilityUnlocked,
			Sensitive:                true,
			Language:                 "en",
			CreatedWithApplicationID: "6b0cd164-8497-4cd5-bec9-957886fac5df",
			VisibilityAdvanced: &gtsmodel.VisibilityAdvanced{
				Federated: true,
				Boostable: true,
				Replyable: false,
				Likeable:  false,
			},
			ActivityStreamsType: gtsmodel.ActivityStreamsNote,
		},
	}
}

// NewTestTags returns a map of gts model tags keyed by their name
func NewTestTags() map[string]*gtsmodel.Tag {
	return map[string]*gtsmodel.Tag{
		"welcome": {
			ID:                     "a7e8f5ca-88a1-4652-8079-a187eab8d56e",
			Name:                   "welcome",
			FirstSeenFromAccountID: "",
			CreatedAt:              time.Now().Add(-71 * time.Hour),
			UpdatedAt:              time.Now().Add(-71 * time.Hour),
			Useable:                true,
			Listable:               true,
			LastStatusAt:           time.Now().Add(-71 * time.Hour),
		},
	}
}

// NewTestFaves returns a map of gts model faves, keyed in the format [faving_account]_[target_status]
func NewTestFaves() map[string]*gtsmodel.StatusFave {
	return map[string]*gtsmodel.StatusFave{
		"local_account_1_admin_account_status_1": {
			ID:              "fc4d42ef-631c-4125-bd9d-88695131284c",
			CreatedAt:       time.Now().Add(-47 * time.Hour),
			AccountID:       "580072df-4d03-4684-a412-89fd6f7d77e6", // local account 1
			TargetAccountID: "8020dbb4-1e7b-4d99-a872-4cf94e64210f", // admin account
			StatusID:        "502ccd6f-0edf-48d7-9016-2dfa4d3714cd", // admin account status 1
		},
	}
}