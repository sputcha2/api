package tests

import (
	"encoding/json"
	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/common/datastore"
	"github.com/HackIllinois/api/services/registration/config"
	"github.com/HackIllinois/api/services/registration/models"
	"github.com/HackIllinois/api/services/registration/service"
	"reflect"
	"testing"
)

var db database.Database

func init() {
	db_connection, err := database.InitDatabase(config.REGISTRATION_DB_HOST, config.REGISTRATION_DB_NAME)

	if err != nil {
		panic(err)
	}

	db = db_connection
}

/*
	Initialize db with test user and mentor info
*/
func SetupTestDB(t *testing.T) {
	user_registration := getBaseUserRegistration()
	err := db.Insert("attendees", &user_registration)

	if err != nil {
		t.Fatal(err)
	}

	mentor_registration := getBaseMentorRegistration()
	err = db.Insert("mentors", &mentor_registration)

	if err != nil {
		t.Fatal(err)
	}
}

/*
	Drop test db
*/
func CleanupTestDB(t *testing.T) {
	err := db.DropDatabase()

	if err != nil {
		t.Fatal(err)
	}
}

/*
	Service level test for getting user registration from db
*/
func TestGetUserRegistrationService(t *testing.T) {
	SetupTestDB(t)

	user_registration, err := service.GetUserRegistration("testid")

	if err != nil {
		t.Fatal(err)
	}

	expected_registration := getBaseUserRegistration()

	if !reflect.DeepEqual(user_registration.Data["firstName"], expected_registration.Data["firstName"]) {
		t.Errorf("Wrong user info.\nExpected %v\ngot %v\n", expected_registration.Data["firstName"], user_registration.Data["firstName"])
	}

	CleanupTestDB(t)
}

/*
	Service level test for creating user registration in the db
*/
func TestCreateUserRegistrationService(t *testing.T) {
	SetupTestDB(t)

	new_registration := getBaseUserRegistration()
	new_registration.Data["id"] = "testid2"
	new_registration.Data["firstName"] = "first2"
	new_registration.Data["lastName"] = "last2"
	err := service.CreateUserRegistration("testid2", new_registration)

	if err != nil {
		t.Fatal(err)
	}

	user_registration, err := service.GetUserRegistration("testid2")

	if err != nil {
		t.Fatal(err)
	}

	expected_registration := getBaseUserRegistration()
	expected_registration.Data["id"] = "testid2"
	expected_registration.Data["firstName"] = "first2"
	expected_registration.Data["lastName"] = "last2"

	if !reflect.DeepEqual(user_registration.Data["firstName"], expected_registration.Data["firstName"]) {
		t.Errorf("Wrong user info.\nExpected %v\ngot %v\n", expected_registration.Data["firstName"], user_registration.Data["firstName"])
	}

	CleanupTestDB(t)
}

/*
	Service level test for updating user registration in the db
*/
func TestUpdateUserRegistrationService(t *testing.T) {
	SetupTestDB(t)

	updated_registration := getBaseUserRegistration()
	updated_registration.Data["id"] = "testid"
	updated_registration.Data["firstName"] = "first2"
	updated_registration.Data["lastName"] = "last2"
	err := service.UpdateUserRegistration("testid", updated_registration)

	if err != nil {
		t.Fatal(err)
	}

	user_registration, err := service.GetUserRegistration("testid")

	if err != nil {
		t.Fatal(err)
	}

	expected_registration := getBaseUserRegistration()
	expected_registration.Data["id"] = "testid"
	expected_registration.Data["firstName"] = "first2"
	expected_registration.Data["lastName"] = "last2"

	if !reflect.DeepEqual(user_registration.Data["firstName"], expected_registration.Data["firstName"]) {
		t.Errorf("Wrong user info.\nExpected %v\ngot %v\n", expected_registration.Data["firstName"], user_registration.Data["firstName"])
	}

	CleanupTestDB(t)
}

/*
	Service level test for getting mentor registration from db
*/
func TestGetMentorRegistrationService(t *testing.T) {
	SetupTestDB(t)

	mentor_registration, err := service.GetMentorRegistration("testid")

	if err != nil {
		t.Fatal(err)
	}

	expected_registration := getBaseMentorRegistration()

	if !reflect.DeepEqual(mentor_registration.Data["firstName"], expected_registration.Data["firstName"]) {
		t.Errorf("Wrong mentor info.\nExpected %v\ngot %v\n", expected_registration.Data["firstName"], mentor_registration.Data["firstName"])
	}

	CleanupTestDB(t)
}

/*
	Service level test for creating mentor registration in the db
*/
func TestCreateMentorRegistrationService(t *testing.T) {
	SetupTestDB(t)

	new_registration := getBaseMentorRegistration()
	new_registration.Data["id"] = "testid2"
	new_registration.Data["firstName"] = "first2"
	new_registration.Data["lastName"] = "last2"
	err := service.CreateMentorRegistration("testid2", new_registration)

	if err != nil {
		t.Fatal(err)
	}

	mentor_registration, err := service.GetMentorRegistration("testid2")

	if err != nil {
		t.Fatal(err)
	}

	expected_registration := getBaseMentorRegistration()
	expected_registration.Data["id"] = "testid2"
	expected_registration.Data["firstName"] = "first2"
	expected_registration.Data["lastName"] = "last2"

	if !reflect.DeepEqual(mentor_registration.Data["firstName"], expected_registration.Data["firstName"]) {
		t.Errorf("Wrong mentor info.\nExpected %v\ngot %v\n", expected_registration.Data["firstName"], mentor_registration.Data["firstName"])
	}

	CleanupTestDB(t)
}

/*
	Service level test for updating mentor registration in the db
*/
func TestUpdateMentorRegistrationService(t *testing.T) {
	SetupTestDB(t)

	updated_registration := getBaseMentorRegistration()
	updated_registration.Data["id"] = "testid"
	updated_registration.Data["firstName"] = "first2"
	updated_registration.Data["lastName"] = "last2"
	err := service.UpdateMentorRegistration("testid", updated_registration)

	if err != nil {
		t.Fatal(err)
	}

	mentor_registration, err := service.GetMentorRegistration("testid")

	if err != nil {
		t.Fatal(err)
	}

	expected_registration := getBaseMentorRegistration()
	expected_registration.Data["id"] = "testid"
	expected_registration.Data["firstName"] = "first2"
	expected_registration.Data["lastName"] = "last2"

	if !reflect.DeepEqual(mentor_registration.Data["firstName"], expected_registration.Data["firstName"]) {
		t.Errorf("Wrong mentor info.\nExpected %v\ngot %v\n", expected_registration.Data["firstName"], mentor_registration.Data["firstName"])
	}

	CleanupTestDB(t)
}

/*
	Service level test for filtering user registrations in the db
*/
func TestGetFilteredUserRegistrationsService(t *testing.T) {
	SetupTestDB(t)

	registration_1 := getBaseUserRegistration()

	registration_2 := getBaseUserRegistration()
	registration_2.Data["id"] = "testid2"

	err := service.CreateUserRegistration(registration_2.Data["id"].(string), registration_2)
	if err != nil {
		t.Fatal(err)
	}

	// Test single value and one keys
	parameters := map[string][]string{
		"id": []string{"testid"},
	}
	user_registrations, err := service.GetFilteredUserRegistrations(parameters)
	if err != nil {
		t.Fatal(err)
	}

	expected_registrations := models.FilteredRegistrations{
		[]models.UserRegistration{
			registration_1,
		},
	}

	if len(user_registrations.Registrations) != len(expected_registrations.Registrations) {
		t.Errorf("Wrong number of registrations.\nExpected %v\ngot %v\n", len(expected_registrations.Registrations), len(user_registrations.Registrations))
	}

	if !reflect.DeepEqual(user_registrations.Registrations[0].Data["firstName"], expected_registrations.Registrations[0].Data["firstName"]) {
		t.Errorf("Wrong user info.\nExpected %v\ngot %v\n", expected_registrations.Registrations[0].Data["firstName"], user_registrations.Registrations[0].Data["firstName"])
	}

	// Test multiple values
	parameters = map[string][]string{
		"id": []string{"testid,testid2"},
	}
	user_registrations, err = service.GetFilteredUserRegistrations(parameters)
	if err != nil {
		t.Fatal(err)
	}

	expected_registrations = models.FilteredRegistrations{
		[]models.UserRegistration{
			registration_1,
			registration_2,
		},
	}

	if len(user_registrations.Registrations) != len(expected_registrations.Registrations) {
		t.Errorf("Wrong number of registrations.\nExpected %v\ngot %v\n", len(expected_registrations.Registrations), len(user_registrations.Registrations))
	}

	if !reflect.DeepEqual(user_registrations.Registrations[1].Data["firstName"], expected_registrations.Registrations[1].Data["firstName"]) {
		t.Errorf("Wrong user info.\nExpected %v\ngot %v\n", expected_registrations.Registrations[1].Data["firstName"], user_registrations.Registrations[1].Data["firstName"])
	}

	// Test type casting
	parameters = map[string][]string{
		"firstName": []string{"first"},
		"age":       []string{"20"},
		"isNovice":  []string{"true"},
	}
	user_registrations, err = service.GetFilteredUserRegistrations(parameters)
	if err != nil {
		t.Fatal(err)
	}

	expected_registrations = models.FilteredRegistrations{
		[]models.UserRegistration{
			registration_1,
			registration_2,
		},
	}

	if len(user_registrations.Registrations) != len(expected_registrations.Registrations) {
		t.Errorf("Wrong number of registrations.\nExpected %v\ngot %v\n", len(expected_registrations.Registrations), len(user_registrations.Registrations))
	}

	if !reflect.DeepEqual(user_registrations.Registrations[1].Data["firstName"], expected_registrations.Registrations[1].Data["firstName"]) {
		t.Errorf("Wrong user info.\nExpected %v\ngot %v\n", expected_registrations.Registrations[1].Data["firstName"], user_registrations.Registrations[1].Data["firstName"])
	}

	CleanupTestDB(t)
}

/*
	Returns a basic user registration
*/
func getBaseUserRegistration() datastore.DataStore {
	base_user_registration := datastore.NewDataStore(config.REGISTRATION_DEFINITION)
	json.Unmarshal([]byte(user_registration_data), &base_user_registration)
	return base_user_registration
}

/*
	Returns a basic mentor registration
*/
func getBaseMentorRegistration() datastore.DataStore {
	base_mentor_registration := datastore.NewDataStore(config.MENTOR_REGISTRATION_DEFINITION)
	json.Unmarshal([]byte(user_registration_data), &base_mentor_registration)
	return base_mentor_registration
}

var user_registration_data string = `
{
	"id": "testid",
	"firstName": "first",
	"lastName": "last",
	"email": "test@gmail.com",
	"shirtSize": "M",
	"diet": "NONE",
	"age": 20,
	"graduationYear": 2020,
	"transportation": "NONE",
	"school": "University of Illinois at Urbana-Champaign",
	"major": "Computer Science",
	"gender": "MALE",
	"professionalInterest": "INTERNSHIP",
	"github": "githubusername",
	"linkedin": "linkedinusername",
	"interests": "Software",
	"isNovice": true,
	"isPrivate": false,
	"phoneNumber": "555-928-7402",
	"longforms": [
		{
			"response": "This is a longform."
		}
	],
	"extraInfos": [
		{
			"response": "This is an extra info."
		}
	],
	"osContributors": [
		{
			"name": "Person",
			"contactInfo": "person@gmail.com"
		}
	],
	"collaborators": [
		{
			"github": "persongithub"
		}
	],
	"createdAt": 10,
	"updatedAt": 15
}
`

var mentor_registration_data string = `
{
	"id": "testid",
	"firstName": "first",
	"lastName": "last",
	"email": "test@gmail.com",
	"shirtSize": "M",
	"github": "githubusername",
	"linkedin": "linkedinusername",
	"createdAt": 10,
	"updatedAt": 15
}
`
