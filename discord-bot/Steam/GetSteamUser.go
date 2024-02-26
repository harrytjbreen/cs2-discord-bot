type ResolveResponse struct {
	Response struct {
		SteamID string `json:"steamid"`
	} `json:"response"`
}



func GetSteamUser(profile string) nil {
    //TODO implement get basic user info
}

func SteamUserExists(profile string) bool {
    userId, err = getSteamUserID(profile)
    if(err != nil) {
        return false
    }


    return true
}

func getSteamUserID(profile string) string, error {
    // Regular expressions for different formats
	steamID32Regex := regexp.MustCompile(`^STEAM_0:(0|1):([0-9]+)$`)
	steamID64Regex := regexp.MustCompile(`^([0-9]{17})$`)
	steamID3Regex := regexp.MustCompile(`^\[U:1:([0-9]+)\]$`)
	profileURLRegex := regexp.MustCompile(`steamcommunity.com/(profiles|id)/([^/]+)/?$`)

	// Check if input matches any of the formats
	if steamID32Regex.MatchString(input) {
		parts := steamID32Regex.FindStringSubmatch(input)
		// Calculate SteamID64 from SteamID32
		authBit, _ := strconv.Atoi(parts[2])
		userBit, _ := strconv.Atoi(parts[3])
		steamID64 := strconv.FormatInt(int64(userBit*2+authBit), 10)
		return steamID64, nil
	} else if steamID64Regex.MatchString(input) {
		return input, nil
	} else if steamID3Regex.MatchString(input) {
		parts := steamID3Regex.FindStringSubmatch(input)
		steamID64 := strconv.FormatInt(76561197960265728+int64(2*userBit), 10)
		return steamID64, nil
	} else if profileURLRegex.MatchString(input) {
		parts := profileURLRegex.FindStringSubmatch(input)
		// Assuming custom URL or SteamID64 in the URL
		customURL := parts[len(parts)-1]
		// If it's a custom URL, we need to resolve it
		if !strings.HasPrefix(customURL, "765") {
			steamID64, err := resolveCustomURL(customURL)
			if err != nil {
				return "", err
			}
			return steamID64, nil
		}
		return customURL, nil
	}

	// If none of the formats match, return an error
	return "", errors.New("invalid input format")
}

func resolveCustomURL(customURL string) (string, error) {
	// Make a request to the Steam Web API
	apiKey := GetEnvVar("STEAM_API_KEY")
	resp, err := http.Get(fmt.Sprintf("http://api.steampowered.com/ISteamUser/ResolveVanityURL/v0001/?key=%sY&vanityurl=%s", apiKey, customURL))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Decode the response
	var resolveResp ResolveResponse
	if err := json.NewDecoder(resp.Body).Decode(&resolveResp); err != nil {
		return "", err
	}

	// Check if the response contains a valid SteamID
	if resolveResp.Response.SteamID == "" {
		return "", fmt.Errorf("unable to resolve custom URL: %s", customURL)
	}

	return resolveResp.Response.SteamID, nil
}
