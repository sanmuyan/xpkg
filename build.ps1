$currentTag = git describe --tags

if ($currentTag -match "v(\d+)\.(\d+)\.(\d+)") {
    $major = [int]$matches[1]
    $minor = [int]$matches[2]
    $patch = [int]$matches[3]

    if ($patch -eq 99 ) {
        $minor = $minor + 1
        $patch = -1
    }

    if ($minor -eq 100) {
        $major = $major + 1
        $minor = 0
    }

    $newPatch = $patch + 1

    $newTag = "v$major.$minor.$newPatch"
    Write-Output "New tag: $newTag"
    git add .
    git commit -m $newTag
    git tag $newTag
    git push
    git push --tags
} else {
    Write-Error "Tag format is not valid. Expected format: vMAJOR.MINOR.PATCH"
}
