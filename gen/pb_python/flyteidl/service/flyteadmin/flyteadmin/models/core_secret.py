# coding: utf-8

"""
    flyteidl/service/admin.proto

    No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)  # noqa: E501

    OpenAPI spec version: version not set
    
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


import pprint
import re  # noqa: F401

import six

from flyteadmin.models.secret_mount_type import SecretMountType  # noqa: F401,E501


class CoreSecret(object):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    """
    Attributes:
      swagger_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    swagger_types = {
        'name': 'str',
        'mount_requirement': 'SecretMountType'
    }

    attribute_map = {
        'name': 'name',
        'mount_requirement': 'mount_requirement'
    }

    def __init__(self, name=None, mount_requirement=None):  # noqa: E501
        """CoreSecret - a model defined in Swagger"""  # noqa: E501

        self._name = None
        self._mount_requirement = None
        self.discriminator = None

        if name is not None:
            self.name = name
        if mount_requirement is not None:
            self.mount_requirement = mount_requirement

    @property
    def name(self):
        """Gets the name of this CoreSecret.  # noqa: E501

        The name of the secret to mount. This has to match an existing secret in the system. It's up to the implementation of the secret management system to require case sensitivity.  # noqa: E501

        :return: The name of this CoreSecret.  # noqa: E501
        :rtype: str
        """
        return self._name

    @name.setter
    def name(self, name):
        """Sets the name of this CoreSecret.

        The name of the secret to mount. This has to match an existing secret in the system. It's up to the implementation of the secret management system to require case sensitivity.  # noqa: E501

        :param name: The name of this CoreSecret.  # noqa: E501
        :type: str
        """

        self._name = name

    @property
    def mount_requirement(self):
        """Gets the mount_requirement of this CoreSecret.  # noqa: E501


        :return: The mount_requirement of this CoreSecret.  # noqa: E501
        :rtype: SecretMountType
        """
        return self._mount_requirement

    @mount_requirement.setter
    def mount_requirement(self, mount_requirement):
        """Sets the mount_requirement of this CoreSecret.


        :param mount_requirement: The mount_requirement of this CoreSecret.  # noqa: E501
        :type: SecretMountType
        """

        self._mount_requirement = mount_requirement

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.swagger_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value
        if issubclass(CoreSecret, dict):
            for key, value in self.items():
                result[key] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, CoreSecret):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
